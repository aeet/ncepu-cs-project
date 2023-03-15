import { Component, OnInit } from '@angular/core';
import { STColumn } from '@yelon/abc/st';
import { SFSchema } from '@yelon/form';
import { _HttpClient } from '@yelon/theme';

@Component({
  selector: 'app-campus-list',
  template: ` <app-simple [schema]="schema" [title]="title" [path]="path" [columns]="columns"></app-simple> `,
  styles: []
})
export class CampusListComponent {
  schema: SFSchema = {
    properties: {
      name: {
        type: 'string',
        ui: {
          i18n: 'name',
          widget: 'string'
        }
      },
      address: {
        type: 'string',
        ui: {
          i18n: 'address',
          widget: 'string'
        }
      }
    }
  };
  title: string = '校区管理';
  path: string = 'campus';
  columns: STColumn[] = [
    { title: '名称', type: '', index: 'name' },
    { title: '地址', type: '', index: 'address' }
  ];
}
