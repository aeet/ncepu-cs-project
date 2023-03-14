import { Injectable, OnDestroy } from '@angular/core';
import { environment } from '@env/environment';
import { _HttpClient } from '@yelon/theme';
import { Observable, Subject } from 'rxjs';

@Injectable({ providedIn: 'root' })
export class RoleService implements OnDestroy {
  private destroy$ = new Subject<void>();
  constructor(private http: _HttpClient) {}

  query(): Observable<any> {
    return this.http.get(`${environment['path']}/role`);
  }

  ngOnDestroy(): void {
    this.destroy$.complete();
  }
}
