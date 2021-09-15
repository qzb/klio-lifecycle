#!/usr/bin/env node

import path from 'path';
import schemaParser from '@apidevtools/json-schema-ref-parser';
import { compile } from 'json-schema-to-typescript';

const schemaPath = path.resolve(process.argv[2]);
const schema = await schemaParser.dereference(schemaPath);
const types = await compile(schema);

console.log(types);
