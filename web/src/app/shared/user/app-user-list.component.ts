import { Component, EventEmitter, OnInit, Output, ViewChild } from '@angular/core';
import { OnReuseInit, ReuseHookOnReuseInitType } from '@yelon/abc/reuse-tab';
import { STColumn, STColumnButton } from '@yelon/abc/st';
import { SFSchema } from '@yelon/form';
import { NzModalService } from 'ng-zorro-antd/modal';

import { SimpleComponent } from '../simple/simple.component';
import { SimpleService } from '../simple/simple.service';

@Component({
  selector: `users`,
  template: `
    <app-simple
      #simple
      (onUserSelect)="userSelect($event)"
      [buttons]="buttons"
      [schema]="schema"
      [title]="title"
      [path]="path"
      [columns]="columns"
    ></app-simple>
  `
})
export class UsersComponent implements OnReuseInit {
  _onReuseInit(type?: ReuseHookOnReuseInitType | undefined): void {
    this.simple?.ngOnInit();
  }
  @ViewChild('simple', { static: false }) simple?: SimpleComponent;
  schema: SFSchema = {
    properties: {
      username: {
        type: 'string',
        title: '姓名',
        ui: {
          widget: 'string'
        }
      },
      account: {
        type: 'string',
        title: '账号',
        ui: {
          widget: 'string'
        }
      },
      passwd: {
        type: 'string',
        title: '密码',
        ui: {
          widget: 'string'
        }
      },
      email: {
        type: 'string',
        title: '邮箱',
        ui: {
          widget: 'string'
        }
      }
    }
  };
  title: string = '用户管理';
  path: string = 'user';
  columns: STColumn[] = [
    { title: '姓名', type: '', index: 'username' },
    { title: '账号', type: '', index: 'account' },
    { title: '邮箱', type: '', index: 'email' }
  ];
  buttons: STColumnButton[] = [];

  constructor(private service: SimpleService<any>, private modal: NzModalService) {
    this.service.init({ path: this.path, title: this.title, schema: this.schema });
  }

  @Output() onUserSelect: EventEmitter<any> = new EventEmitter<any>();

  userSelect(users: any): void {
    this.onUserSelect.emit(users);
  }
}
