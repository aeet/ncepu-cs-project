import { Component, OnInit, ViewChild } from '@angular/core';
import { OnReuseInit, ReuseHookOnReuseInitType } from '@yelon/abc/reuse-tab';
import { STColumn } from '@yelon/abc/st';
import { SFSchema } from '@yelon/form';
import { _HttpClient } from '@yelon/theme';
import { SimpleComponent } from 'src/app/shared/simple/simple.component';

@Component({
  selector: 'app-campus-list',
  template: ` <app-simple #simple [schema]="schema" [title]="title" [path]="path" [columns]="columns"></app-simple> `,
  styles: []
})
export class CampusListComponent implements OnReuseInit {
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
