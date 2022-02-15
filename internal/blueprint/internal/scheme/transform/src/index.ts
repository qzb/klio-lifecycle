import type * as v1beta4 from './types/g2a-cli/v1beta4';
import type * as current from './types/g2a-cli/v2.0';
import type * as internal from './types/internal';

export function toCurrent(obj: v1beta4.Object | current.Object): current.Object {
  if (obj.apiVersion === 'g2a-cli/v2.0') {
    return obj;
  }

  if (obj.kind === 'Project') {
    return toCurrentProject(obj);
  }

  if (obj.kind === 'Service') {
    return toCurrentService(obj);
  }

  if (obj.kind === 'Environment') {
    return toCurrentEnvironment(obj);
  }

  return obj; // fails unless all kinds are handled
}

export function toInternal(obj: v1beta4.Object | current.Object): internal.Object {
  if (obj.apiVersion === 'g2a-cli/v1beta4') {
    return toInternal(toCurrent(obj));
  }

  if (obj.kind === 'Project') {
    return toInternalProject(obj);
  }

  if (obj.kind === 'Service') {
    return toInternalService(obj);
  }

  if (obj.kind === 'Environment') {
    return toInternalEnvironment(obj);
  }

  if (
    obj.kind === 'Builder' ||
    obj.kind === 'Deployer' ||
    obj.kind === 'Pusher' ||
    obj.kind === 'Tagger'
  ) {
    return toInternalExecutor(obj);
  }

  return obj; // fails unless all kinds are handled
}

function toCurrentProject(obj: v1beta4.Project): Required<current.Project> {
  const files = [
    {
      git: {
        url: 'git@github.com:g2a-com/klio-lifecycle.git',
        rev: 'main',
        files: 'assets/executors/*/*.yaml',
      },
    },
    ...(obj.services ?? ['services/*']).map((s) => s.replace(/\/?$/, '/service.yaml')),
    ...(obj.environments ?? ['environments/*']).map((s) => s.replace(/\/?$/, '/environment.yaml')),
  ];

  return {
    apiVersion: 'g2a-cli/v2.0',
    kind: 'Project',
    name: obj.name || 'project',
    files,
    tasks: {},
    variables: {},
  };
}

function toCurrentService(obj: v1beta4.Service): Required<current.Service> {
  const artifacts: current.Service['artifacts'] = obj?.build?.artifacts || [];

  if (obj.hooks?.['pre-build']) {
    artifacts.unshift({
      script: {
        sh: 'set -e\n' + obj.hooks['pre-build'].join('\n'),
      },
      push: false,
    });
  }

  if (obj.hooks?.['post-build']) {
    artifacts.push({
      script: {
        sh: 'set -e\n' + obj.hooks['post-build'].join('\n'),
      },
      push: false,
    });
  }

  const releases = obj?.deploy?.releases || [];

  if (obj.hooks?.['pre-deploy']) {
    releases.unshift({
      script: {
        sh: 'set -e\n' + obj.hooks['pre-deploy'].join('\n'),
      },
    });
  }

  if (obj.hooks?.['post-deploy']) {
    releases.push({
      script: {
        sh: 'set -e\n' + obj.hooks['post-deploy'].join('\n'),
      },
    });
  }

  const tags = Object.entries(obj?.build?.tagPolicy || {}).map(([k, v]) => ({ [k]: v }));

  return {
    apiVersion: 'g2a-cli/v2.0',
    kind: 'Service',
    name: obj.name,
    tags,
    artifacts,
    releases,
    tasks: {},
  };
}

function toCurrentEnvironment(obj: v1beta4.Environment): Required<current.Environment> {
  return {
    apiVersion: 'g2a-cli/v2.0',
    kind: 'Environment',
    name: obj.name,
    deployServices: obj.deployServices || [],
    variables: obj.variables || {},
  };
}

function toInternalProject(obj: current.Project): internal.Project {
  return {
    kind: 'Project',
    files: [obj.files || []]
      .flat()
      .map((f) =>
        typeof f === 'string'
          ? { glob: f }
          : [f.git.files].flat().map((g) => ({ glob: g, git: { url: f.git.url, rev: f.git.rev } }))
      )
      .flat(),
    name: obj.name,
    variables: obj.variables || {},
  };
}

function toInternalService(obj: current.Service): internal.Service {
  return {
    kind: 'Service',
    name: obj.name,
    build: {
      artifacts: {
        toBuild: (obj.artifacts || []).map(toInternalEntry),
        toPush: (obj.artifacts || [])
          .map((e) => {
            if (typeof e === 'string' || !('push' in e)) {
              return e;
            } else {
              return e.push as string | false | Record<string, unknown>;
            }
          })
          .map(toInternalEntry)
          .filter((e) => e.type),
      },
      tags: (obj.tags || []).map(toInternalEntry),
    },
    deploy: {
      releases: (obj.releases || []).map(toInternalEntry),
    },
    run: {
      tasks: Object.entries(obj.tasks || {})
        .map(
          ([k, v]) => [k, v?.map(toInternalEntry)] as [string, ReturnType<typeof toInternalEntry>[]]
        )
        .reduce((a, [k, v]) => {
          a[k] = v;
          return a;
        }, {} as Record<string, ReturnType<typeof toInternalEntry>[]>),
    },
  };
}

function toInternalEnvironment(obj: current.Environment): internal.Environment {
  return {
    kind: 'Environment',
    name: obj.name,
    deployServices: obj.deployServices || [],
    variables: obj.variables || {},
  };
}

function toInternalExecutor(
  obj: current.Tagger | current.Builder | current.Pusher | current.Deployer
): internal.Executor {
  return {
    kind: obj.kind,
    name: obj.name,
    js: obj.js,
    schema: JSON.stringify(obj.schema ?? {}),
  };
}

function toInternalEntry(obj: string | false | Record<string, unknown>, index: number) {
  if (obj === false) {
    return { index, type: '', spec: null };
  }
  if (typeof obj === 'string') {
    return { index, type: obj, spec: null };
  }
  const [type, spec] = Object.entries(obj).filter(kv => kv[0] !== 'push')[0];
  return { index, type, spec };
}
