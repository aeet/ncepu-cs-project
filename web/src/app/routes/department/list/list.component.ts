import { Component, ViewChild } from '@angular/core';
import { OnReuseInit, ReuseHookOnReuseInitType } from '@yelon/abc/reuse-tab';
import { STColumn } from '@yelon/abc/st';
import { SFSchema } from '@yelon/form';
import { _HttpClient } from '@yelon/theme';
import { SimpleComponent } from 'src/app/shared/simple/simple.component';

@Component({
  selector: 'app-department-list',
  template: ` <app-simple #simple [schema]="schema" [title]="title" [path]="path" [columns]="columns"></app-simple> `,
  styles: []
})
export class DepartmentListComponent implements OnReuseInit {
  _onReuseInit(type?: ReuseHookOnReuseInitType | undefined): void {
    this.simple?.ngOnInit();
  }
  @ViewChild('simple', { static: false }) simple?: SimpleComponent;

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
