package schema

import (
	"bytes"
	"fmt"
	"io"
	"path"
	"strings"

	"github.com/g2a-com/cicd/internal/placeholders"
	"gopkg.in/yaml.v3"
)

const LatestVersion = "g2a-cli/v2.0"

// Migrator updates documents to newer versions.
type Migrator interface {
	// Migrate updates content of a YAML or a JSON file. Keep in mind that input
	// should be pre-validated.
	Migrate([]byte) ([]byte, error)
}

// MigrationError records errors
type MigrationError struct {
	Node    *yaml.Node
	Message string
}

func (e *MigrationError) Error() string {
	return fmt.Sprintf("error in line %d: %s", e.Node.Line, e.Message)
}

// Migrate updates content of a YAML or a JSON file to the latest version. Keep
// in mind that input should be pre-validated.
func Migrate(input []byte) ([]byte, error) {
	return NewMigrator(LatestVersion).Migrate(input)
}

type migrator struct {
	TargetVersion string
}

// NewMigrator creates new migrator which updates documents to specified
// version.
func NewMigrator(targetVersion string) Migrator {
	return &migrator{targetVersion}
}

func (m *migrator) Migrate(input []byte) ([]byte, error) {
	result := &bytes.Buffer{}
	decoder := yaml.NewDecoder(bytes.NewReader(input))
	encoder := yaml.NewEncoder(result)

	for {
		var document yaml.Node

		err := decoder.Decode(&document)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		err = m.migrateDocument(&document)
		if err != nil {
			return nil, err
		}

		err = encoder.Encode(&document)
		if err != nil {
			return nil, err
		}
	}

	return result.Bytes(), nil
}

func (m *migrator) migrateDocument(rootNode *yaml.Node) error {
	if rootNode.Kind == yaml.DocumentNode {
		return m.migrateDocument(rootNode.Content[0])
	}

	versionNode := findMapValue(rootNode, "apiVersion")
	if versionNode == nil {
		return &MigrationError{rootNode, "missing apiVersion field"}
	}

	version := versionNode.Value
	if version == m.TargetVersion || version == LatestVersion {
		return nil
	}

	err := migrateDocumentToNextVersion(rootNode)
	if err != nil {
		return err
	}

	return m.migrateDocument(rootNode)
}

// migrateDocumentToNextVersion migrates document to subsequent version. If
// document is in latest version returns MigrationError.
func migrateDocumentToNextVersion(rootNode *yaml.Node) error {
	versionNode := findMapValue(rootNode, "apiVersion")
	if versionNode == nil {
		return &MigrationError{rootNode, "missing apiVersion field"}
	}

	switch versionNode.Value {
	case "g2a-cli/v1beta4":
		err := migrateDocumentFromV1Beta4ToV2(rootNode)
		if err != nil {
			return err
		}
	default:
		return &MigrationError{versionNode, fmt.Sprintf("unsupported version: %s", versionNode.Value)}
	}

	return nil
}

// migrateDocumentFromV1Beta4ToV2 migrates any document using v1beta4 version to v2.0.
func migrateDocumentFromV1Beta4ToV2(rootNode *yaml.Node) (err error) {
	// Update apiVersion
	findMapValue(rootNode, "apiVersion").Value = "g2a-cli/v2.0"

	// Migrate document content based on its kind.
	switch kind := findMapValue(rootNode, "kind").Value; kind {
	case "Project":
		err = migrateProjectFromV1Beta4ToV2(rootNode)
	case "Service":
		err = migrateServiceFromV1Beta4ToV2(rootNode)
	case "Environment":
		err = migrateEnvironmentFromV1Beta4ToV2(rootNode)
	default:
		err = &MigrationError{rootNode, fmt.Sprintf("unsupported kind: %s", kind)}
	}
	if err != nil {
		return err
	}

	// Fix order of apiVersion, kind and name. Since fields are prepended, they
	// are processed in reversed order.
	for _, field := range []string{"name", "kind", "apiVersion"} {
		if idx := findMapKeyIndex(rootNode, field); idx != -1 {
			kv := rootNode.Content[idx : idx+2]
			rootNode.Content = spliceNodes(rootNode.Content, idx, 2)
			rootNode.Content = spliceNodes(rootNode.Content, 0, 0, kv...)
		}
	}

	// Migrate placeholders names
	migratePlaceholderNamesFromV1Beta4ToV2(rootNode)

	return
}

// toV2Project migrates Project from v1beta4 to v2.0.
func migrateProjectFromV1Beta4ToV2(rootNode *yaml.Node) error {
	// Check for property name conflicts.
	for _, name := range []string{"files", "tasks", "variables"} {
		if findMapKeyIndex(rootNode, name) != -1 {
			return &MigrationError{rootNode, fmt.Sprintf("cannot migrate project to v2.0, document already contains property %q", name)}
		}
	}

	// Add empty "files" field, it's content is populated in the furter part of
	// this function.
	filesNode := createSequenceNode()
	rootNode.Content = append(rootNode.Content, createStrScalarNode("files"), filesNode)

	// FIXME: replace this remote with builtin executors.
	filesNode.Content = append(filesNode.Content, createMappingNode(
		createStrScalarNode("git"),
		createMappingNode(
			createStrScalarNode("url"),
			createStrScalarNode("git@github.com:g2a-com/klio-lifecycle.git"),
			createStrScalarNode("rev"),
			createStrScalarNode("main"),
			createStrScalarNode("files"),
			createStrScalarNode("assets/executors/*/*.yaml"),
		),
	))

	// Add name if there is none.
	if idx := findMapKeyIndex(rootNode, "name"); idx == -1 {
		// Position of the "name" field is fixed by parent function, so put it at
		// the end for now.
		rootNode.Content = append(rootNode.Content, createStrScalarNode("name"), createStrScalarNode("project"))
	}

	// Migrate services to files
	if idx := findMapKeyIndex(rootNode, "services"); idx != -1 {
		// Each entry is appended with "service.yaml" name and added to "files".
		for _, node := range rootNode.Content[idx+1].Content {
			node.Value = path.Join(node.Value, "service.yaml")
			filesNode.Content = append(filesNode.Content, node)
		}
		// Remove "services" field
		rootNode.Content = spliceNodes(rootNode.Content, idx, 2)
	} else {
		// If there is no "services" field, use default value (v2.0 has no defaults)
		filesNode.Content = append(filesNode.Content, createStrScalarNode("services/*/service.yaml"))
	}

	// Migrate environments to files
	if idx := findMapKeyIndex(rootNode, "environments"); idx != -1 {
		// Each entry is appended with "environment.yaml" name and added to "files".
		for _, node := range rootNode.Content[idx+1].Content {
			node.Value = path.Join(node.Value, "environment.yaml")
			filesNode.Content = append(filesNode.Content, node)
		}
		// Remove "environments" field
		rootNode.Content = spliceNodes(rootNode.Content, idx, 2)
	} else {
		// If there is no "environments" field, use default value (v2.0 has no defaults)
		filesNode.Content = append(filesNode.Content, createStrScalarNode("environments/*/environment.yaml"))
	}

	return nil
}

// toV2Service migrates Service from v1beta4 to v2.0.
func migrateServiceFromV1Beta4ToV2(rootNode *yaml.Node) error {
	// Check for property name conflicts.
	for _, name := range []string{"artifacts", "tags", "releases", "tasks"} {
		if findMapKeyIndex(rootNode, name) != -1 {
			return &MigrationError{rootNode, fmt.Sprintf("cannot migrate service to v2.0, document already contains property %q", name)}
		}
	}

	// Find and remove "hooks", "build", and "deploy" fields. Those fields are not
	// present in v2.0. Theirs content is migrated in further part of this
	// function.
	hooksNode := createMappingNode()
	if idx := findMapKeyIndex(rootNode, "hooks"); idx != -1 {
		hooksNode = rootNode.Content[idx+1]
		rootNode.Content = spliceNodes(rootNode.Content, idx, 2)
	}
	buildNode := createMappingNode()
	if idx := findMapKeyIndex(rootNode, "build"); idx != -1 {
		buildNode = rootNode.Content[idx+1]
		rootNode.Content = spliceNodes(rootNode.Content, idx, 2)
	}
	deployNode := createMappingNode()
	if idx := findMapKeyIndex(rootNode, "deploy"); idx != -1 {
		deployNode = rootNode.Content[idx+1]
		rootNode.Content = spliceNodes(rootNode.Content, idx, 2)
	}

	// Check for unsupported hooks
	for _, stage := range []string{"push", "publish", "lint", "test"} {
		for _, prefix := range []string{"pre-", "post-"} {
			if findMapValue(hooksNode, prefix+stage) != nil {
				return &MigrationError{hooksNode, fmt.Sprintf("cannot migrate %q hook", prefix+stage)}
			}
		}
	}

	// Properties in a service should have following order:
	//   - apiVersion, kind, name - those are used by all document kinds
	//   - tags,
	//   - artifacts,
	//   - releases,
	//   - ...other properties in the same order as in the original document.
	// To enforce this order this function prepends "releases", "artifacts" and
	// "tags". We don't need to worry about order of "apiVersion", "kind"
	// and "name", since it's fixed by parent function afterwards.

	// Migrate deploy.releases. It's mostly unchanged, but since v2.0 doesn't
	// support hooks, we need to convert them to releases.
	if releasesNode := findMapValue(deployNode, "releases"); releasesNode != nil {
		// Migrate hooks
		if hookNode := findMapValue(hooksNode, "pre-deploy"); hookNode != nil {
			releasesNode.Content = spliceNodes(releasesNode.Content, 0, 0, hookToEntry(hookNode))
		}
		if hookNode := findMapValue(hooksNode, "post-deploy"); hookNode != nil {
			releasesNode.Content = append(releasesNode.Content, hookToEntry(hookNode))
		}

		// Preprend to the document
		rootNode.Content = spliceNodes(rootNode.Content, 0, 0, createStrScalarNode("releases"), releasesNode)
	}

	// Migrate build.artifacts. It's mostly unchanged, but since v2.0 doesn't
	// support hooks, we need to convert them to artifacts.
	if artifactsNode := findMapValue(buildNode, "artifacts"); artifactsNode != nil {
		// Migrate hooks
		if hookNode := findMapValue(hooksNode, "pre-build"); hookNode != nil {
			entry := hookToEntry(hookNode)
			entry.Content = append(entry.Content, createStrScalarNode("push"), createBoolScalarNode(false))
			artifactsNode.Content = spliceNodes(artifactsNode.Content, 0, 0, entry)
		}
		if hookNode := findMapValue(hooksNode, "post-build"); hookNode != nil {
			entry := hookToEntry(hookNode)
			entry.Content = append(entry.Content, createStrScalarNode("push"), createBoolScalarNode(false))
			artifactsNode.Content = append(artifactsNode.Content, entry)
		}

		// Preprend to the document
		rootNode.Content = spliceNodes(rootNode.Content, 0, 0, createStrScalarNode("artifacts"), artifactsNode)
	}

	// Migrate build.tagPolicy. This field was replaced with "tags" field, which
	// follows the same format as "artifacts" and "releases" ("tagPolicy" was a
	// map, "tags" is a list of maps).
	if tagPolicyNode := findMapValue(buildNode, "tagPolicy"); tagPolicyNode != nil {
		tagsNode := createSequenceNode()

		// Copy comments from the tagPolicy node
		tagsNode.HeadComment = tagPolicyNode.HeadComment
		tagsNode.LineComment = tagPolicyNode.LineComment
		tagsNode.FootComment = tagPolicyNode.FootComment

		// Copy key/value items from tagPolicy to separate maps in tags list
		for i := 0; i+1 < len(tagPolicyNode.Content); i += 2 {
			tagsNode.Content = append(tagsNode.Content, createMappingNode(tagPolicyNode.Content[i:i+2]...))
		}

		// Prepend to the document
		rootNode.Content = spliceNodes(rootNode.Content, 0, 0, createStrScalarNode("tags"), tagsNode)
	}

	return nil
}

// toV2Environment migrates Environment from v1beta4 to v2.0.
func migrateEnvironmentFromV1Beta4ToV2(rootNode *yaml.Node) error {
	// There are no changes between v1beta4 and v2.0
	return nil
}

// hookToEntry converts list of shell commands to service entry using "script"
// executor.
func hookToEntry(node *yaml.Node) *yaml.Node {
	builder := strings.Builder{}
	builder.WriteString("set -e\n")

	for _, cmd := range node.Content {
		builder.WriteString(cmd.Value)
		builder.WriteString("\n")
	}

	return createMappingNode(
		createStrScalarNode("script"),
		createMappingNode(
			createStrScalarNode("sh"),
			createStrScalarNode(builder.String()),
		),
	)
}

// migratePlaceholderNamesFromV1Beta4ToV2 replaces legacy placeholders with new ones.
func migratePlaceholderNamesFromV1Beta4ToV2(node *yaml.Node) {
	if node.Tag != "!!str" {
		for _, n := range node.Content {
			migratePlaceholderNamesFromV1Beta4ToV2(n)
		}
		return
	}

	value, err := placeholders.Replace(node.Value, func(name string) (result string, err error) {
		id := strings.ToLower(name)
		switch {
		case id == ".dirs.project":
			result = "{{ .Project.Dir }}"
		case id == ".dirs.service":
			result = "{{ .Service.Dir }}"
		case id == ".dirs.environment":
			result = "{{ .Environment.Dir }}"
		case id == ".opts.tag":
			result = "{{ .Tag }}"
		case strings.HasPrefix(id, ".env."):
			result = fmt.Sprintf("{{ .Environment.Vars.%s }}", name[5:])
		default:
			result = fmt.Sprintf("{{ %s }}", name)
		}
		return
	})
	// Replace should not return error, but handle it just to be safe.
	if err != nil {
		panic(err)
	}

	node.Value = value.(string)
}

// findMapKeyIndex returns index of a key node within map, value node is
// stored in consecutive item, so it's index is a key index + 1. This function
// supports only string keys.
//
// If there is no key matching provided value, function returns -1.
func findMapKeyIndex(node *yaml.Node, name string) int {
	if node.Kind != yaml.MappingNode {
		panic("document has to be a map")
	}

	for i := 0; i+1 < len(node.Content); i += 2 {
		key := node.Content[i]
		if key.Kind == yaml.ScalarNode && key.Tag == "!!str" && key.Value == name {
			return i
		}
	}

	return -1
}

// findMapValue returns value node for specified key, or nil if it cannot be
// found. This function supports only string keys.
func findMapValue(node *yaml.Node, name string) *yaml.Node {
	if idx := findMapKeyIndex(node, name); idx != -1 {
		return node.Content[idx+1]
	}
	return nil
}

// spliceNodes removes `count` elements starting from `start` and adds
// `addItems` in their place.
func spliceNodes(items []*yaml.Node, start, count int, addItems ...*yaml.Node) (ret []*yaml.Node) {
	ret = make([]*yaml.Node, len(items)-count+len(addItems))
	copy(ret, items[:start])
	copy(ret[start:], addItems)
	copy(ret[start+len(addItems):], items[start+count:])
	return
}

// createMappingNode creates yaml mapping node with provided content.
func createMappingNode(content ...*yaml.Node) *yaml.Node {
	return &yaml.Node{
		Kind:    yaml.MappingNode,
		Tag:     "!!map",
		Content: content,
	}
}

// createMappingNode creates yaml sequence node with provided content.
func createSequenceNode(content ...*yaml.Node) *yaml.Node {
	return &yaml.Node{
		Kind:    yaml.SequenceNode,
		Tag:     "!!seq",
		Content: content,
	}
}

// createMappingNode creates yaml scalar node storing provided string value.
func createStrScalarNode(value string) *yaml.Node {
	return &yaml.Node{
		Kind:  yaml.ScalarNode,
		Tag:   "!!str",
		Value: value,
	}
}

// createMappingNode creates yaml scalar node storing provided boolean value.
func createBoolScalarNode(value bool) *yaml.Node {
	return &yaml.Node{
		Kind:  yaml.ScalarNode,
		Tag:   "!!bool",
		Value: fmt.Sprintf("%v", value),
	}
}
