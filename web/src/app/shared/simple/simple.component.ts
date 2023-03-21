import { Component, EventEmitter, Input, OnDestroy, OnInit, Output, TemplateRef } from '@angular/core';
import { STChange, STColumn, STColumnButton, STColumnSelection, STData } from '@yelon/abc/st';
import { SFSchema } from '@yelon/form';

import { SimpleService } from './simple.service';

@Component({
  selector: `app-simple`,
  template: `
    <page-header [action]="phActionTpl">
      <ng-template #phActionTpl>
        <button *ngIf="!template" (click)="add()" nz-button nzType="primary">新建</button>
        <ng-template [ngTemplateOutlet]="template"></ng-template>
      </ng-template>
    </page-header>
    <nz-card>
      <st #st [data]="data" [res]="{ process: dataProcess }" [loading]="loading" [columns]="columns" (change)="change($event)"></st>
    </nz-card>
  `,
  styles: []
})
export class SimpleComponent implements OnInit {
  @Input() schema: SFSchema = { properties: {} };
  @Input() title: string = '';
  @Input() path: string = '';
  @Input() columns: STColumn[] = [];
  @Input() buttons: STColumnButton[] = [];
  @Input() check: boolean = true;

  @Input() template: TemplateRef<any> | null = null;

  // eslint-disable-next-line @angular-eslint/no-output-on-prefix
  @Output() readonly onUserSelect: EventEmitter<any> = new EventEmitter<any>();
  data: [] = [];
  loading: boolean = false;

  constructor(private service: SimpleService<any>) {}

  ngOnInit(): void {
    this.service.init({ title: this.title, schema: this.schema, path: this.path });
    this.query();
    if (this.check) {
      this.columns = [
        {
          title: '编号',
          index: 'id',
          type: 'checkbox'
        },
        ...this.columns
      ];
    }
    const a = this.columns.filter(col => col.buttons);
    if (a.length === 0) {
      this.columns = [
        ...this.columns,
        {
          title: '',
          buttons: this.buttons.concat([
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
          ])
        }
      ];
    }
  }

  query(): void {
    this.loading = true;
    this.service.query().subscribe((response: any) => {
      this.data = response.body;
      this.loading = false;
    });
  }

  update(role: any): void {
    const id = role.id;
    this.service.createForm(role => {
      role.id = id;
      return this.service
        .update(role)
        .toPromise()
        .then(() => {
          this.query();
        });
    }, role);
  }

  add(): void {
    this.service.createForm(role => {
      return this.service
        .add(role)
        .toPromise()
        .then(() => {
          this.query();
        });
    }, null);
  }

  delete(id: string): void {
    this.loading = true;
    this.service.delete(id).subscribe(() => {
      this.loading = false;
      this.query();
    });
  }

  dataProcess(data: STData[]): STData[] {
    return data.map((i, index) => {
      i.checked = false;
      return i;
    });
  }

  change(e: STChange): void {
    if (e.type === 'checkbox') {
      this.onUserSelect.emit(e.checkbox);
    }
  }
}
