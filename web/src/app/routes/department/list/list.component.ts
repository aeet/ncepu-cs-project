import { Component } from '@angular/core';
import { STColumn } from '@yelon/abc/st';
import { SFSchema } from '@yelon/form';
import { _HttpClient } from '@yelon/theme';

@Component({
  selector: 'app-department-list',
  template: ` <app-simple [schema]="schema" [title]="title" [path]="path" [columns]="columns"></app-simple> `,
  styles: []
})
export class DepartmentListComponent {
  schema: SFSchema = {
    properties: {
      name: {
        type: 'string',
        ui: {
          i18n: 'name',
          widget: 'string'
        }
      },
      code: {
        type: 'string',
        ui: {
          i18n: 'code',
          widget: 'string'
        }
      },
      description: {
        type: 'string',
        ui: {
          i18n: 'description',
          widget: 'string'
        }
      }
    }
  };
  title: string = '系部';
  path: string = 'department';
  columns: STColumn[] = [
    { title: '名称', type: '', index: 'name' },
    { title: '编码', type: '', index: 'code' },
    { title: '描述', type: '', index: 'code' }
  ];
}
