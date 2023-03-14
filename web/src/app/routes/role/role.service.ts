import { Injectable, OnDestroy } from '@angular/core';
import { environment } from '@env/environment';
import { SFComponent, SFSchema, SFValue, SFValueChange } from '@yelon/form';
import { _HttpClient } from '@yelon/theme';
import { NzModalRef, NzModalService } from 'ng-zorro-antd/modal';
import { Observable, Subject, takeUntil } from 'rxjs';

@Injectable({ providedIn: 'root' })
export class RoleService implements OnDestroy {
  private destroy$ = new Subject<void>();
  constructor(private http: _HttpClient, private modal: NzModalService) {}

  createForm(nzOnOk: (param: any) => Promise<any>, data?: any): void {
    const schema: SFSchema = {
      properties: {
        role_name: {
          type: 'string',
          ui: {
            i18n: 'role_name',
            widget: 'string'
          }
        },
        role_value: {
          type: 'string',
          ui: {
            i18n: 'role_value',
            widget: 'string'
          }
        }
      }
    };
    let formValue: SFValue = {};
    const ref: NzModalRef<SFComponent, any> = this.modal.create<SFComponent>({
      nzTitle: '角色管理',
      nzMask: true,
      nzClosable: true,
      nzMaskClosable: true,
      nzContent: SFComponent,
      nzComponentParams: {
        button: null,
        formData: data
      },
      nzOnOk: () => {
        let valid: boolean | undefined = ref.componentInstance?.validator({ emitError: true, onlyRoot: false });
        if (!valid) {
          return Promise.reject(false);
        } else {
          return nzOnOk(formValue);
        }
      }
    });
    ref.componentInstance?.formValueChange.pipe(takeUntil(this.destroy$)).subscribe((sfValueChange: SFValueChange) => {
      formValue = sfValueChange.value;
    });
    ref.componentInstance?.refreshSchema(schema);
    // hack: use reset to trigger formValueChange function
    ref.componentInstance?.reset();
  }

  query(): Observable<any> {
    return this.http.get(`${environment['path']}/role`);
  }

  add(role: any): Observable<any> {
    return this.http.post(`${environment['path']}/role`, role);
  }

  delete(id: string): Observable<any> {
    return this.http.delete(`${environment['path']}/role/${id}`);
  }

  update(role: any): Observable<any> {
    return this.http.put(`${environment['path']}/role`, role);
  }

  ngOnDestroy(): void {
    this.destroy$.complete();
  }
}
