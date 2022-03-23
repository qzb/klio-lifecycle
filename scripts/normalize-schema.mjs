#!/usr/bin/env node

import path from "path";
import schemaParser from "@apidevtools/json-schema-ref-parser";

const schemaPath = path.resolve(process.argv[2]);
const schema = await schemaParser.dereference(schemaPath);

delete schema["$defs"];

console.log(JSON.stringify(schema, null, 2));
