import { Component, Input, ViewChild } from '@angular/core';
import { OnReuseInit, ReuseHookOnReuseInitType } from '@yelon/abc/reuse-tab';
import { STColumn, STColumnButton } from '@yelon/abc/st';
import { SFComponent, SFSchema, SFSelectWidgetSchema, SFValueChange } from '@yelon/form';
import { _HttpClient } from '@yelon/theme';
import { NzModalService } from 'ng-zorro-antd/modal';
import { map } from 'rxjs';
import { SimpleComponent } from 'src/app/shared/simple/simple.component';
import { SimpleService } from 'src/app/shared/simple/simple.service';

@Component({
  selector: 'app-major-list',
  template: ` <app-simple #simple [buttons]="buttons" [schema]="schema" [title]="title" [path]="path" [columns]="columns"></app-simple> `,
  styles: []
})
export class MajorListComponent implements OnReuseInit {
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
      },
      special_type: {
        type: 'string',
        ui: {
          i18n: 'special_type',
          widget: 'string'
        }
      },
      enrollment_type: {
        type: 'string',
        ui: {
          i18n: 'enrollment_type',
          widget: 'string'
        }
      },
      is_major_category: {
        type: 'boolean',
        ui: {
          i18n: 'is_major_category',
          widget: 'boolean'
        }
      },
      major_category: {
        type: 'string',
        ui: {
          i18n: 'major_category',
          widget: 'string'
        }
      }
    }
  };
  title: string = '专业管理';
  path: string = 'major';
  columns: STColumn[] = [
    { title: '名称', type: '', index: 'name' },
    { title: '代码', type: '', index: 'code' },
    { title: '描述', type: '', index: 'description' },
    { title: '特殊类型', type: '', index: 'special_type' },
    { title: '招生类型', type: '', index: 'enrollment_type' },
    { title: '是否专业大类', type: '', index: 'is_major_category' },
    { title: '专业大类', type: '', index: 'major_category' }
  ];
  buttons: STColumnButton[] = [
    {
      text: '关联系部',
      click: major => {
        let value: any = {};
        const schema: SFSchema = {
          properties: {
            department: {
              title: '系部',
              type: 'string',
              ui: {
                widget: 'select',
                asyncData: () =>
                  this.departmentService.query().pipe(
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
            formData: { department: major?.edges?.department?.id },
            button: null
          },
          nzOnOk: () => {
            major = {
              ...major,
              edges: {
                department: {
                  id: value['department']
                }
              }
            };
            return this.majorService
              .update(major)
              .toPromise()
              .then(() => {
                this._onReuseInit();
              });
          }
        });
        ref.componentInstance?.formValueChange.subscribe((sfValueChange: SFValueChange) => {
          value = sfValueChange.value;
        });
        ref.componentInstance?.refreshSchema(schema);
        // hack: use reset to trigger formValueChange function
        ref.componentInstance?.reset();
        this.majorService.init({ schema: this.schema, title: this.title, path: this.path });
      }
    }
  ];

  get departmentService(): SimpleService<any> {
    this.service.init({ path: 'department' });
    return this.service;
  }

  get majorService(): SimpleService<any> {
    this.service.init({ path: 'major' });
    return this.service;
  }

  constructor(private service: SimpleService<any>, private modal: NzModalService) {}
}
