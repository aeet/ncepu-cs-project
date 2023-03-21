import { Component, OnInit } from '@angular/core';
import { environment } from '@env/environment';
import { OnReuseInit, ReuseHookOnReuseInitType } from '@yelon/abc/reuse-tab';
import { STColumn, STComponent } from '@yelon/abc/st';
import { SFSchema, SFSelectWidgetSchema } from '@yelon/form';
import { _HttpClient } from '@yelon/theme';
import { NzModalService } from 'ng-zorro-antd/modal';
import { map } from 'rxjs';

@Component({
  selector: `app-student-query`,
  template: `
    <page-header [action]="phActionTpl">
      <ng-template #phActionTpl> </ng-template>
      <sf [schema]="querySchema" layout="inline" (formSubmit)="search($event)" (formReset)="reset()"></sf>
    </page-header>
    <nz-card>
      <st #st [columns]="columns" [data]="data"></st>
    </nz-card>
  `
})
export class StudentQueryComponent implements OnReuseInit, OnInit {
  columns: STColumn[] = [
    { title: '姓名', index: 'name', type: '' },
    { title: '年龄', index: 'age', type: '' },
    { title: '性别', index: 'sex', type: '' },
    { title: '身份证号', index: 'code', type: '' },
    {
      type: '',
      title: '操作',
      width: '100px',
      buttons: [
        {
          text: '查看实践经历',
          click: s => {
            const r = this.modal.create({
              nzTitle: '实践信息',
              nzWidth: 1200,
              nzContent: STComponent,
              nzComponentParams: {
                data: s.edges.practical_experience
              }
            });
            r.componentInstance?.resetColumns({
              columns: [
                { title: '名称', index: 'name', type: '' },
                { title: '实践单位', index: 'unit', type: '' },
                { title: '开始时间', index: 'start_time', type: '' },
                { title: '结束时间', index: 'end_time', type: '' },
                { title: '实践描述', index: 'string', type: '' }
              ]
            });
            console.log(s);
          }
        },
        {
          text: '查看家庭信息',
          click: s => {
            const r = this.modal.create({
              nzTitle: '家庭信息',
              nzWidth: 1200,
              nzContent: STComponent,
              nzComponentParams: {
                data: s.edges.family_info
              }
            });
            r.componentInstance?.resetColumns({
              columns: [
                { title: '家庭成员姓名', index: 'name', type: '' },
                { title: '与学生关系', index: 'relationship', type: '' },
                { title: '身份证号', index: 'id_card', type: '' },
                { title: '年龄', index: 'age', type: '' },
                { title: '职业', index: 'occupation', type: '' },
                { title: '职务', index: 'post', type: '' },
                { title: '工作单位', index: 'work_unit', type: '' },
                { title: '联系电话', index: 'contact_number', type: '' },
                { title: '健康状况', index: 'health', type: '' }
              ]
            });
            console.log(s);
          }
        },
        {
          text: '查看证书',
          click: s => {
            const r = this.modal.create({
              nzTitle: '证书信息',
              nzWidth: 1200,
              nzContent: STComponent,
              nzComponentParams: {
                data: s.edges.certificate
              }
            });
            r.componentInstance?.resetColumns({
              columns: [
                { title: '名称', index: 'name', type: '' },
                { title: '证书编号', index: 'code', type: '' },
                { title: '证书描述', index: 'description', type: '' },
                { title: '颁发机构', index: 'department', type: '' },
                { title: '颁发日期', index: 'issue_date', type: '' },
                { title: '证书类型', index: 'certificate_type', type: '' },
                { title: '证书级别', index: 'certificate_level', type: '' },
                { title: '证书类别', index: 'certificate_type2', type: '' },
                { title: '证书类别细分', index: 'award_category', type: '' }
              ]
            });
          }
        }
      ]
    }
  ];
  data: any = [];
  temp_data: any = [];
  querySchema: SFSchema = {
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
      } as SFSelectWidgetSchema,
      department: {
        title: '系部',
        type: 'string',
        ui: {
          widget: 'select',
          asyncData: () =>
            this.http.get(`${environment['path']}/department`).pipe(
              map((response: any) => {
                return response.body.map((item: any) => {
                  return { label: item.name, value: item.id };
                });
              })
            )
        }
      } as SFSelectWidgetSchema,
      major: {
        title: '专业',
        type: 'string',
        ui: {
          widget: 'select',
          asyncData: () =>
            this.http.get(`${environment['path']}/major`).pipe(
              map((response: any) => {
                return response.body.map((item: any) => {
                  return { label: item.name, value: item.id };
                });
              })
            )
        }
      } as SFSelectWidgetSchema,
      class: {
        title: '班级',
        type: 'string',
        ui: {
          widget: 'select',
          asyncData: () =>
            this.http.get(`${environment['path']}/class`).pipe(
              map((response: any) => {
                return response.body.map((item: any) => {
                  return { label: item.name, value: item.id };
                });
              })
            )
        }
      } as SFSelectWidgetSchema,
      name: {
        title: '名称',
        type: 'string',
        ui: {
          widget: 'string'
        }
      } as SFSelectWidgetSchema
    }
  };
  constructor(private http: _HttpClient, private modal: NzModalService) {}
  ngOnInit(): void {
    this.query(null);
  }
  _onReuseInit(type?: ReuseHookOnReuseInitType | undefined): void {
    this.query(null);
  }

  query(search?: any): void {
    this.http.get(`${environment['path']}/student`, search).subscribe(response => {
      this.data = response.body;
      this.temp_data = response.body;
    });
  }

  search(value: any): void {
    console.log(value);
    if (value.department) {
      this.data = this.data.filter((item: any) => item.edges.department.id === value.department);
    }
    if (value.major) {
      this.data = this.data.filter((item: any) => item.edges.major.id === value.major);
    }
    if (value.class) {
      this.data = this.data.filter((item: any) => item.edges.class.id === value.class);
    }
    if (value.name) {
      this.data = this.data.filter((item: any) => item.name.includes(value.name));
    }
  }
  reset(): void {
    this.data = this.temp_data;
  }
}
