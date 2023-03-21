import { Component, OnInit, ViewChild } from '@angular/core';
import { environment } from '@env/environment';
import { OnReuseInit, ReuseHookOnReuseInitType } from '@yelon/abc/reuse-tab';
import { STColumn, STColumnButton } from '@yelon/abc/st';
import { SFComponent, SFDateWidgetSchema, SFSchema, SFSelectWidgetSchema, SFValueChange } from '@yelon/form';
import { _HttpClient } from '@yelon/theme';
import { NzModalService } from 'ng-zorro-antd/modal';
import { map } from 'rxjs';
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
  buttons: STColumnButton[] = [
    {
      text: '添加实践经历',
      click: z => {
        let value: any = {};
        const schema: SFSchema = {
          properties: {
            name: {
              type: 'string',
              title: '实践名称',
              ui: {
                widget: 'string'
              }
            },
            unit: {
              type: 'string',
              title: '实践单位',
              ui: {
                widget: 'string'
              }
            },
            start_time: {
              type: 'string',
              title: '开始时间',
              ui: {
                widget: 'date',
                mode: 'date'
              } as SFDateWidgetSchema
            },
            end_time: {
              type: 'string',
              title: '结束时间',
              ui: {
                widget: 'date',
                mode: 'date'
              } as SFDateWidgetSchema
            },
            describe: {
              type: 'string',
              title: '实践描述',
              ui: {
                widget: 'string'
              }
            }
          }
        };
        const ref = this.modal.create<SFComponent>({
          nzTitle: '添加实践经历',
          nzContent: SFComponent,
          nzComponentParams: {
            schema: {},
            formData: {},
            button: null
          },
          nzOnOk: () => {
            value = {
              ...value,
              edges: {
                student: {
                  id: z.id
                }
              }
            };
            return this.http.post(`${environment['path']}/student/practical`, value).toPromise();
          }
        });
        ref.componentInstance?.formValueChange.subscribe((sfValueChange: SFValueChange) => {
          value = sfValueChange.value;
        });
        ref.componentInstance?.refreshSchema(schema);
        // hack: use reset to trigger formValueChange function
        ref.componentInstance?.reset();
      }
    },
    {
      text: '添加家庭成员',
      click: z => {
        let value: any = {};
        const schema: SFSchema = {
          properties: {
            name: {
              type: 'string',
              title: '家庭成员姓名',
              ui: {
                widget: 'string'
              }
            },
            relationship: {
              type: 'string',
              title: '与学生关系',
              ui: {
                widget: 'string'
              }
            },
            id_card: {
              type: 'string',
              title: '身份证号',
              ui: {
                widget: 'string'
              }
            },
            age: {
              type: 'string',
              title: '年龄',
              ui: {
                widget: 'string'
              }
            },
            occupation: {
              type: 'string',
              title: '职业',
              ui: {
                widget: 'string'
              }
            },
            post: {
              type: 'string',
              title: '职务',
              ui: {
                widget: 'string'
              }
            },
            work_unit: {
              type: 'string',
              title: '工作单位',
              ui: {
                widget: 'string'
              }
            },
            contact_number: {
              type: 'string',
              title: '联系电话',
              ui: {
                widget: 'string'
              }
            },
            health: {
              type: 'string',
              title: '健康状况',
              ui: {
                widget: 'string'
              }
            }
          }
        };
        const ref = this.modal.create<SFComponent>({
          nzTitle: '添加家庭成员',
          nzContent: SFComponent,
          nzComponentParams: {
            schema: {},
            formData: {},
            button: null
          },
          nzOnOk: () => {
            value = {
              ...value,
              edges: {
                student: {
                  id: z.id
                }
              }
            };
            return this.http.post(`${environment['path']}/student/family`, value).toPromise();
          }
        });
        ref.componentInstance?.formValueChange.subscribe((sfValueChange: SFValueChange) => {
          value = sfValueChange.value;
        });
        ref.componentInstance?.refreshSchema(schema);
        // hack: use reset to trigger formValueChange function
        ref.componentInstance?.reset();
      }
    },
    {
      text: '添加证书信息',
      click: z => {
        let value: any = {};
        const schema: SFSchema = {
          properties: {
            name: {
              type: 'string',
              title: '证书名称',
              ui: {
                widget: 'string'
              }
            },
            code: {
              type: 'string',
              title: '证书编号',
              ui: {
                widget: 'string'
              }
            },
            description: {
              type: 'string',
              title: '证书描述',
              ui: {
                widget: 'string'
              }
            },
            department: {
              type: 'string',
              title: '颁发机构',
              ui: {
                widget: 'string'
              }
            },
            issue_date: {
              type: 'string',
              title: '颁发日期',
              ui: {
                widget: 'date',
                mode: 'date'
              } as SFDateWidgetSchema
            },
            certificate_type: {
              type: 'string',
              title: '证书类型',
              ui: {
                widget: 'string'
              }
            },
            certificate_level: {
              type: 'string',
              title: '证书级别',
              ui: {
                widget: 'string'
              }
            },
            certificate_type2: {
              type: 'string',
              title: '证书类别',
              ui: {
                widget: 'string'
              }
            },
            award_category: {
              type: 'string',
              title: '证书类别细分',
              ui: {
                widget: 'string'
              }
            }
          }
        };
        const ref = this.modal.create<SFComponent>({
          nzTitle: '添加证书',
          nzContent: SFComponent,
          nzComponentParams: {
            schema: {},
            formData: {},
            button: null
          },
          nzOnOk: () => {
            value = {
              ...value,
              edges: {
                student: {
                  id: z.id
                }
              }
            };
            return this.http.post(`${environment['path']}/student/certificate`, value).toPromise();
          }
        });
        ref.componentInstance?.formValueChange.subscribe((sfValueChange: SFValueChange) => {
          value = sfValueChange.value;
        });
        ref.componentInstance?.refreshSchema(schema);
        // hack: use reset to trigger formValueChange function
        ref.componentInstance?.reset();
      }
    },
    {
      text: '关联班级',
      click: s => {
        let value: any = {};
        const schema: SFSchema = {
          properties: {
            class: {
              title: '班级',
              type: 'string',
              ui: {
                widget: 'select',
                asyncData: () =>
                  this.http.get(`${environment['path']}/class`).pipe(
                    map(response => {
                      return response.body.map((item: any) => {
                        return { label: item.name, value: item.id };
                      });
                    })
                  )
              }
            } as SFSelectWidgetSchema
          }
        };
        const ref = this.modal.create<SFComponent>({
          nzTitle: '关联系部',
          nzContent: SFComponent,
          nzComponentParams: {
            schema: { properties: {} },
            formData: { department: s?.edges?.class?.id },
            button: null
          },
          nzOnOk: () => {
            value = {
              ...s,
              edges: {
                class: {
                  id: value['class']
                }
              }
            };
            return this.http.put(`${environment['path']}/student`, value).toPromise();
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
