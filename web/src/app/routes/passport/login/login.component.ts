import { ChangeDetectionStrategy, ChangeDetectorRef, Component, Inject, OnDestroy, Optional } from '@angular/core';
import { FormBuilder, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { StartupService } from '@core';
import { ReuseTabService } from '@yelon/abc/reuse-tab';
import { YA_SERVICE_TOKEN, ITokenService, SocialService } from '@yelon/auth';
import { SettingsService, _HttpClient } from '@yelon/theme';
import { NzTabChangeEvent } from 'ng-zorro-antd/tabs';

@Component({
  selector: 'passport-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.less'],
  providers: [SocialService],
  changeDetection: ChangeDetectionStrategy.OnPush
})
export class UserLoginComponent implements OnDestroy {
  constructor(private fb: FormBuilder) {}

  form = this.fb.nonNullable.group({
    username: ['', [Validators.required, Validators.pattern(/^(admin|user)$/)]],
    password: ['', [Validators.required, Validators.pattern(/^(ng\-yunzai\.com)$/)]]
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
  }

  ngOnDestroy(): void {
    if (this.interval$) {
      clearInterval(this.interval$);
    }
  }
}
