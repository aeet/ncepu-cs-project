import { Injectable, OnDestroy } from '@angular/core';
import { environment } from '@env/environment';
import { SFComponent, SFSchema, SFValue, SFValueChange } from '@yelon/form';
import { _HttpClient } from '@yelon/theme';
import { NzModalRef, NzModalService } from 'ng-zorro-antd/modal';
import { takeUntil, Observable, Subject } from 'rxjs';

export interface Simple {
  title: string;
  path: string;
  schema: SFSchema;
}

@Injectable({ providedIn: 'root' })
export class SimpleService<T extends Simple> implements OnDestroy {
  private destroy$ = new Subject<void>();

  t: T = {} as T;

  constructor(private http: _HttpClient, private modal: NzModalService) {}

  init(t: T) {
    this.t = {} as T;
    this.t = t;
  }

  createForm(nzOnOk: (param: any) => Promise<any>, data?: any): void {
    const schema: SFSchema = this.t.schema;
    let formValue: SFValue = {};
    const ref: NzModalRef<SFComponent, any> = this.modal.create<SFComponent>({
      nzTitle: this.t.title,
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
    return this.http.get(`${environment['path']}/${this.t.path}`);
  }

  add(role: any): Observable<any> {
    return this.http.post(`${environment['path']}/${this.t.path}`, role);
  }

  delete(id: string): Observable<any> {
    return this.http.delete(`${environment['path']}/${this.t.path}/${id}`);
  }

  update(role: any): Observable<any> {
    return this.http.put(`${environment['path']}/${this.t.path}`, role);
  }

  ngOnDestroy(): void {
    this.destroy$.complete();
  }
}
