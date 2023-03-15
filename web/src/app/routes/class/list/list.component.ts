import { Component } from '@angular/core';
import { STColumn } from '@yelon/abc/st';
import { SFSchema } from '@yelon/form';
import { _HttpClient } from '@yelon/theme';

@Component({
  selector: 'app-class-list',
  template: ` <app-simple [schema]="schema" [title]="title" [path]="path" [columns]="columns"></app-simple> `,
  styles: []
})
export class ClassListComponent {
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
      },
      type: {
        type: 'string',
        ui: {
          i18n: 'type',
          widget: 'string'
        }
      }
    }
  };
  title: string = '班级管理';
  path: string = 'class';
  columns: STColumn[] = [
    { title: '名称', type: '', index: 'name' },
    { title: '编号', type: '', index: 'code' },
    { title: '描述', type: '', index: 'description' },
    { title: '类型', type: '', index: 'type' }
  ];
}
