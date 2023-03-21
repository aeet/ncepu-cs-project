import { Component, OnInit, ViewChild } from '@angular/core';
import { environment } from '@env/environment';
import { OnReuseInit, ReuseHookOnReuseInitType } from '@yelon/abc/reuse-tab';
import { STColumn, STColumnButton } from '@yelon/abc/st';
import { SFSchema } from '@yelon/form';
import { _HttpClient } from '@yelon/theme';
import { NzModalService } from 'ng-zorro-antd/modal';
import { SimpleComponent } from 'src/app/shared/simple/simple.component';
import { SimpleService } from 'src/app/shared/simple/simple.service';
import { UsersComponent } from 'src/app/shared/user/app-user-list.component';

@Component({
  selector: `student-list`,
  template: `
    <ng-template #btn>
      <button (click)="studentAdd()" nz-button nzType="primary">新建</button>
    </ng-template>
    <app-simple
      #simple
      [template]="btn"
      [buttons]="buttons"
      [schema]="schema"
      [title]="title"
      [path]="path"
      [columns]="columns"
    ></app-simple>
  `
})
export class StudentListComponent implements OnReuseInit {
  _onReuseInit(type?: ReuseHookOnReuseInitType | undefined): void {
    this.setInit();
  }

  setInit(): void {
    this.service.init({ path: this.path, title: this.title, schema: this.schema });
  }

  @ViewChild('simple', { static: false }) simple?: SimpleComponent;
  schema: SFSchema = {
    properties: {
      name: {
        type: 'string',
        title: '姓名',
        ui: {
          widget: 'string'
        }
      },
      sex: {
        type: 'string',
        title: '性别',
        ui: {
          widget: 'string'
        }
      },
      age: {
        type: 'number',
        title: '年龄',
        ui: {
          widget: 'number'
        }
      },
      code: {
        type: 'string',
        title: '身份证号',
        ui: {
          widget: 'string'
        }
      }
    }
  };
  title: string = '学生信息管理';
  path: string = 'student';
  columns: STColumn[] = [
    { title: '姓名', index: 'name', type: '' },
    { title: '年龄', index: 'age', type: '' },
    { title: '性别', index: 'sex', type: '' },
    { title: '身份证号', index: 'code', type: '' }
  ];
  buttons: STColumnButton[] = [];

  constructor(private modal: NzModalService, private http: _HttpClient, private service: SimpleService<any>) {
    this.setInit();
  }
  studentAdd(): void {
    let us: any = [];
    this.modal.create({
      nzTitle: '添加学生',
      nzContent: UsersComponent,
      nzWidth: 800,
      nzComponentParams: {
        userSelect: users => {
          us = users;
        }
      },
      nzOnOk: () =>
        this.http
          .post(`${environment['path']}/student/batch`, us)
          .toPromise()
          .then(() => {
            this.setInit();
            this.simple?.query();
          })
    });
  }
}
