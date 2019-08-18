import { Application } from 'typedoc';

const app = new Application({
  includeDeclarations: true,
});
// app.generateJson(['./angularjs/index.d.ts'], './docs/docs.json');
app.generateDocs(['./angularjs/index.d.ts'], './docs');
