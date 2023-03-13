import { ChangeDetectorRef, Component, Inject, OnDestroy, Optional } from '@angular/core';
import { FormBuilder, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { StartupService } from '@core';
import { UserService } from '@shared';
import { ReuseTabService } from '@yelon/abc/reuse-tab';
import { ITokenService, SocialService, YA_SERVICE_TOKEN } from '@yelon/auth';
import { SettingsService, _HttpClient } from '@yelon/theme';
import { NzMessageService } from 'ng-zorro-antd/message';
import { NzTabChangeEvent } from 'ng-zorro-antd/tabs';
import { catchError } from 'rxjs';

@Component({
  selector: 'passport-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.less']
})
export class UserLoginComponent implements OnDestroy {
  constructor(
    private fb: FormBuilder,
    private router: Router,
    @Optional()
    @Inject(ReuseTabService)
    private reuseTabService: ReuseTabService,
    @Inject(YA_SERVICE_TOKEN) private tokenService: ITokenService,
    private startupSrv: StartupService,
    private userService: UserService,
    private msg: NzMessageService
  ) {}

  form = this.fb.nonNullable.group({
    username: ['default', [Validators.required]],
    password: ['default', [Validators.required]]
  });
  error = '';
  type = 0;
  loading = false;

  count = 0;
  interval$: any;
  switch({ index }: NzTabChangeEvent): void {
    this.type = index!;
  }

  submit(): void {
    this.error = '';
    const { username, password } = this.form.controls;
    username.markAsDirty();
    username.updateValueAndValidity();
    password.markAsDirty();
    password.updateValueAndValidity();
    if (username.invalid || password.invalid) {
      return;
    }
    this.loading = true;
    this.userService
      .login(username.value, password.value)
      .pipe(
        catchError(e => {
          if (e && e.error && e.error.error) {
            this.msg.error(e.error.error);
          }
          this.loading = false;
          throw e;
        })
      )
      .subscribe((response: any) => {
        this.tokenService.set({ token: response.access_token, expired: response.expires_in * 1000, ...response });
        this.userService.user().subscribe(response => {
          console.log(response);
        });
        this.reuseTabService.clear();
        this.startupSrv.load().subscribe(() => {
          let url = this.tokenService.referrer!.url || '/';
          if (url.includes('/passport')) {
            url = '/';
          }
          this.router.navigateByUrl(url);
        });
        this.loading = false;
      });
  }

  ngOnDestroy(): void {
    if (this.interval$) {
      clearInterval(this.interval$);
    }
  }
}
