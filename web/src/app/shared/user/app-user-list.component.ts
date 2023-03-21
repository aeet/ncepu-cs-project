import { HttpClient } from '@angular/common/http';
import { Component, EventEmitter, OnInit, Output, ViewChild } from '@angular/core';
import { environment } from '@env/environment';
import { OnReuseInit, ReuseHookOnReuseInitType } from '@yelon/abc/reuse-tab';
import { STColumn, STColumnButton } from '@yelon/abc/st';
import { SFComponent, SFSchema, SFSelectWidgetSchema, SFValueChange } from '@yelon/form';
import { NzModalService } from 'ng-zorro-antd/modal';
import { map } from 'rxjs';

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
  buttons: STColumnButton[] = [
    {
      text: '分配角色',
      click: u => {
        let value: any = {};
        const schema: SFSchema = {
          properties: {
            role: {
              title: '角色',
              type: 'string',
              ui: {
                widget: 'select',
                asyncData: () =>
                  this.http.get(`${environment['path']}/role`).pipe(
                    map((response: any) => {
                      return response.body.map((item: any) => {
                        return { label: item.role_name, value: item.id };
                      });
                    })
                  )
              }
            } as SFSelectWidgetSchema
          }
        };
        const ref = this.modal.create<SFComponent>({
          nzTitle: '关联角色',
          nzContent: SFComponent,
          nzComponentParams: {
            schema: { properties: {} },
            formData: { department: u?.edges?.role?.id },
            button: null
          },
          nzOnOk: () => {
            value = {
              ...u,
              edges: {
                role: {
                  id: value['role']
                }
              }
            };
            return this.http.put(`${environment['path']}/user`, value).toPromise();
          }
        });
        ref.componentInstance?.formValueChange.subscribe((sfValueChange: SFValueChange) => {
          value = sfValueChange.value;
        });
        ref.componentInstance?.refreshSchema(schema);
        // hack: use reset to trigger formValueChange function
        ref.componentInstance?.reset();
      }
    }
  ];

  constructor(private service: SimpleService<any>, private modal: NzModalService, private http: HttpClient) {
    this.service.init({ path: this.path, title: this.title, schema: this.schema });
  }

  @Output() onUserSelect: EventEmitter<any> = new EventEmitter<any>();

  userSelect(users: any): void {
    this.onUserSelect.emit(users);
  }
}
