import { HttpContext } from '@angular/common/http';
import { Injectable, OnDestroy } from '@angular/core';
import { environment } from '@env/environment';
import { ALLOW_ANONYMOUS } from '@yelon/auth';
import { _HttpClient } from '@yelon/theme';
import { Observable, Subject, takeUntil } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class UserService implements OnDestroy {
  private destroy$: Subject<any> = new Subject();

  constructor(private http: _HttpClient) {}

  public login(username: string, password: string): Observable<any> {
    return this.http
      .post(`${environment['path']}/auth/user/login`, { username, password }, null, {
        context: new HttpContext().set(ALLOW_ANONYMOUS, true)
      })
      .pipe(takeUntil(this.destroy$));
  }

  public user(): Observable<any> {
    return this.http.get(`${environment['path']}/user`).pipe(takeUntil(this.destroy$));
  }

  public logout(): Observable<any> {
    return this.http.get(`${environment['path']}/auth/logout`);
  }

  ngOnDestroy(): void {
    this.destroy$.complete();
  }

}
