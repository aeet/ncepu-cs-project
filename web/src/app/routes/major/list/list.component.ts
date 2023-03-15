import { Component, ViewChild } from '@angular/core';
import { OnReuseInit, ReuseHookOnReuseInitType } from '@yelon/abc/reuse-tab';
import { STColumn } from '@yelon/abc/st';
import { SFSchema } from '@yelon/form';
import { _HttpClient } from '@yelon/theme';
import { SimpleComponent } from 'src/app/shared/simple/simple.component';

@Component({
  selector: 'app-major-list',
  template: ` <app-simple #simple [schema]="schema" [title]="title" [path]="path" [columns]="columns"></app-simple> `,
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
}
