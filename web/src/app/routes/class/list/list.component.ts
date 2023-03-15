import { Component, ViewChild } from '@angular/core';
import { OnReuseInit, ReuseHookOnReuseInitType } from '@yelon/abc/reuse-tab';
import { STColumn, STColumnButton } from '@yelon/abc/st';
import { SFComponent, SFSchema, SFSelectWidgetSchema, SFValueChange } from '@yelon/form';
import { _HttpClient } from '@yelon/theme';
import { NzModalService } from 'ng-zorro-antd/modal';
import { map } from 'rxjs';
import { SimpleComponent } from 'src/app/shared/simple/simple.component';
import { SimpleService } from 'src/app/shared/simple/simple.service';

@Component({
  selector: 'app-class-list',
  template: ` <app-simple #simple [buttons]="buttons" [schema]="schema" [title]="title" [path]="path" [columns]="columns"></app-simple> `,
  styles: []
})
export class ClassListComponent implements OnReuseInit {
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

  buttons: STColumnButton[] = [
    {
      text: '关联专业',
      click: cls => {
        let value: any = {};
        const schema: SFSchema = {
          properties: {
            major: {
              title: '专业',
              type: 'string',
              ui: {
                widget: 'select',
                asyncData: () =>
                  this.majorService.query().pipe(
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
          nzTitle: '关联专业',
          nzContent: SFComponent,
          nzComponentParams: {
            schema: { properties: {} },
            formData: { major: cls?.edges?.major?.id },
            button: null
          },
          nzOnOk: () => {
            const class_major = {
              ...cls,
              edges: {
                major: {
                  id: value['major']
                }
              }
            };
            return this.classService
              .update(class_major)
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
        this.classService.init({ schema: this.schema, title: this.title, path: this.path });
      }
    }
  ];

  get classService(): SimpleService<any> {
    this.service.init({ path: 'class' });
    return this.service;
  }

  get majorService(): SimpleService<any> {
    this.service.init({ path: 'major' });
    return this.service;
  }

  constructor(private service: SimpleService<any>, private modal: NzModalService) {}
}
