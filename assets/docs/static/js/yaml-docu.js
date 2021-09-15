import { html, render } from 'https://unpkg.com/lit-html@1.1.2/lit-html.js';
import { unsafeHTML } from 'https://unpkg.com/lit-html@1.1.2/directives/unsafe-html.js';

(async function () {
  const tables = document.querySelectorAll('.yaml-config-table');
  console.log(tables);

  for (let table of tables) {
    let path = table.attributes['path'].value.trim();
    let response = await fetch(path);
    let json = await response.json();

    render(html` ${template(json)} `, table);

    if (location.hash) {
      table.querySelector(location.hash).scrollIntoView();
    }
  }
})();

function* template(schema) {
  let lines = prerenderSchema(schema);
  for (let idx = 0; idx < lines.length; idx++) {
    let line = lines[idx];

    let propCass = 'required key'; //line.isRequired ? 'key required' : 'key';
    let valueClass = line.const ? 'value' : 'example';
    let anchorId = line.ref?.slice(2).replace(/\//g, '-');
    let ident = '\u00A0'.repeat(line.identDepth * 2);
    let descrCol = '';
    let hashCol = '';

    if (line.arrayItemStart) {
      ident = ident.replace(/..$/, '-\u00A0');
    }

    if (line.description) {
      let nextDescrIdx = lines.findIndex((l, i) => l.description && i > idx);
      let rowspan = (nextDescrIdx > -1 ? nextDescrIdx : lines.length) - idx;
      descrCol = html`<td class="comment" rowspan="${rowspan}">
        <span>${unsafeHTML(line.description)}</span>
      </td>`;
      hashCol = html` <td class="comment" rowspan="${rowspan}"><span>#&nbsp;</span></td> `;
    }

    let key = line.key
      ? html`<span class="${propCass}">${ident}${anchor(anchorId, line.key)}: </span>`
      : html`<span class="${propCass}">${ident}</span>`;

    yield html`
      <tr>
        <td>${key}<span class="${valueClass}">${line.value}</span></td>
        ${hashCol} ${descrCol}
      </tr>
    `;
  }
}

function anchor(id, label) {
  return html`<a class="anchor" id="${id}"></a><a class="key" href="#${id}">${label}</a>`;
}

function prerenderSchema(schema, ref = '#/', identDepth = 0, inArray = false) {
  let [subSchemaRef, subSchema] = getSubSchema(schema, ref);

  let description = subSchema.description;
  let value =
    subSchema.default ??
    subSchema.const ??
    subSchema.enum?.[0] ??
    subSchema.examples?.[0] ??
    subSchema.items?.examples;

  let lines = [{ identDepth, description, const: !!subSchema.const }];

  if (value) {
    if (value && subSchema.items?.examples && subSchema.items?.['x-examplesDescriptions']) {
      for (let i = 0; i < subSchema.items.examples.length; i++) {
        let subLines = prerenderJSON(subSchema.items.examples[i], identDepth + 1);
        lines.push({
          ...subLines[0],
          arrayItemStart: true,
          description: subSchema.items['x-examplesDescriptions'][i],
        });
        lines.push(...subLines.slice(1));
      }
    } else if (value && typeof value === 'object' && Object.keys(value).length) {
      lines.push(...prerenderJSON(value, identDepth + 1));
    } else if (value && typeof value === 'string' && value.includes('\n')) {
      lines[0].value = '|';
      lines.push(...value.split('\n').map((l) => ({ identDepth, value: l })));
    } else {
      lines[0].value = stringify(value);
    }
  } else if (subSchema.type === 'object') {
    for (let prop of Object.keys(subSchema.properties || {})) {
      let propRef = subSchemaRef + '/properties/' + prop;
      let isRequired = subSchema.required?.includes(prop);
      let subLines = prerenderSchema(schema, propRef, identDepth + 1);

      lines.push({ ...subLines[0], identDepth, isRequired, key: prop, ref: propRef });
      lines.push(...subLines.slice(1));
    }
    if (subSchema.additionalProperties?.description) {
      let propRef = subSchemaRef + '/additionalProperties/';
      let subLines = prerenderSchema(schema, propRef, identDepth + 1);

      lines.push({ ...subLines[0], identDepth, key: 'example', ref: propRef });
      lines.push(...subLines.slice(1));
    }
  } else if (subSchema.type === 'array' && (subSchema.items?.oneOf || subSchema.items?.anyOf)) {
    for (let idx = 0; idx < subSchema?.items?.oneOf.length; idx++) {
      let subLines = prerenderSchema(
        schema,
        subSchemaRef + '/items/oneOf/' + idx,
        identDepth + 1,
        true
      );
      subLines[0].arrayItemStart = true;
      lines.push(...subLines);
    }
  }

  if (ref === '#/' || (inArray && !description)) {
    lines.shift();
  }

  return lines;
}

function prerenderJSON(data, identDepth = 0) {
  if (Array.isArray(data)) {
    return data.flatMap((v) => {
      let singleLine = !v || typeof v !== 'object' || Object.keys(v).length === 0;
      return singleLine
        ? [{ identDepth, arrayItemStart: true, value: stringify(v) }]
        : prerenderJSON(v, identDepth).map((l, i) => ({
            ...l,
            arrayItemStart: l.arrayItemStart || !i,
          }));
    });
  } else {
    return Object.entries(data).flatMap(([key, v]) => {
      let singleLine = !v || typeof v !== 'object' || Object.keys(v).length === 0;
      return singleLine
        ? [{ identDepth, key, value: stringify(v) }]
        : [{ identDepth, key }, ...prerenderJSON(v, identDepth + (Array.isArray(v) ? 2 : 1))];
    });
  }
}

function stringify(value) {
  if (typeof value !== 'string') {
    return JSON.stringify(value);
  }

  if (value.match(/[^a-zA-Z]/)) {
    return JSON.stringify(value);
  }

  return value;
}

function getSubSchema(schema, ref) {
  let get = (r) =>
    r
      .slice(2)
      .split('/')
      .map(decodeURIComponent)
      .filter(Boolean)
      .reduce((s, prop) => s[prop], schema);

  let subSchema, description, examples;
  for (let r = ref; r; ) {
    ref = r;
    subSchema = get(ref);
    description ??= subSchema.description;
    examples ??= subSchema.examples;

    r =
      subSchema.$ref ||
      (subSchema.oneOf && ref + '/oneOf/0') ||
      (subSchema.anyOf && ref + '/anyOf/0');
  }

  return [ref, { ...subSchema, description, examples }];
}
