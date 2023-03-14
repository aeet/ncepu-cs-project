import { Component, OnInit, ViewChild } from '@angular/core';
import { STColumn, STComponent } from '@yelon/abc/st';
import { ModalHelper, _HttpClient } from '@yelon/theme';

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
      <st #st [data]="data" [columns]="columns"></st>
    </nz-card>
  `
})
export class RoleListComponent implements OnInit {
  data = [];
  @ViewChild('st') private readonly st!: STComponent;
  columns: STColumn[] = [
    { title: '角色名称', type: '', index: 'role_name' },
    { title: '角色值', type: '', index: 'role_value' },
    {
      title: '',
      buttons: [{ text: '编辑' }, { text: '删除' }]
    }
  ];

  constructor(private http: _HttpClient, private modal: ModalHelper, private roleService: RoleService) {}

  ngOnInit(): void {
    this.roleService.query().subscribe((response: any) => {
      this.data = response.body;
    });
  }

  add(): void {}
}
