import { Component, OnInit, ViewChild } from '@angular/core';
import { STColumn, STComponent } from '@yelon/abc/st';
import { _HttpClient } from '@yelon/theme';

import { RoleService } from '../role.service';

@Component({
  selector: 'app-role-list',
  template: `
    <page-header [action]="phActionTpl">
      <ng-template #phActionTpl>
        <button (click)="add()" nz-button nzType="primary">新建</button>
      </ng-template>
    </page-header>
    <nz-card>
      <st #st [data]="data" [loading]="loading" [columns]="columns"></st>
    </nz-card>
  `
})
export class RoleListComponent implements OnInit {
  data = [];
  loading = false;
  @ViewChild('st') private readonly st!: STComponent;
  columns: STColumn[] = [
    { title: '角色名称', type: '', index: 'role_name' },
    { title: '角色值', type: '', index: 'role_value' },
    {
      title: '',
      buttons: [
        {
          text: '编辑',
          click: role => {
            this.update(role);
          }
        },
        {
          text: '删除',
          click: role => {
            this.delete(role.id);
          }
        }
      ]
    }
  ];

  constructor(private roleService: RoleService) {}

  ngOnInit(): void {
    this.query();
  }

  query(): void {
    this.loading = true;
    this.roleService.query().subscribe((response: any) => {
      this.data = response.body;
      this.loading = false;
    });
  }

  update(role: any): void {
    const id = role.id;
    this.roleService.createForm(role => {
      role.id = id;
      return this.roleService
        .update(role)
        .toPromise()
        .then(() => {
          this.query();
        });
    }, role);
  }

  add(): void {
    this.roleService.createForm(role => {
      return this.roleService
        .add(role)
        .toPromise()
        .then(() => {
          this.query();
        });
    }, null);
  }

  delete(id: string): void {
    this.loading = true;
    this.roleService.delete(id).subscribe(() => {
      this.loading = false;
      this.query();
    });
  }
}
