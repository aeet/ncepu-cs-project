import { Injectable, Inject } from '@angular/core';
import { UserService } from '@shared';
import { ACLService } from '@yelon/acl';
import { YA_SERVICE_TOKEN, ITokenService } from '@yelon/auth';
import { YUNZAI_I18N_TOKEN, MenuService, SettingsService, TitleService } from '@yelon/theme';
import type { NzSafeAny } from 'ng-zorro-antd/core/types';
import { NzIconService } from 'ng-zorro-antd/icon';
import { Observable, of, map, mergeMap } from 'rxjs';

import { ICONS } from '../../../style-icons';
import { ICONS_AUTO } from '../../../style-icons-auto';
import { I18NService } from '../i18n/i18n.service';

/**
 * Used for application startup
 * Generally used to get the basic data of the application, like: Menu Data, User Data, etc.
 */
@Injectable()
export class StartupService {
  constructor(
    iconSrv: NzIconService,
    private menuService: MenuService,
    @Inject(YUNZAI_I18N_TOKEN) private i18n: I18NService,
    private settingService: SettingsService,
    private aclService: ACLService,
    private titleService: TitleService,
    @Inject(YA_SERVICE_TOKEN) private tokenService: ITokenService,
    private userService: UserService
  ) {
    iconSrv.addIcon(...ICONS_AUTO, ...ICONS);
  }

  private viaMockI18n(): Observable<void> {
    const defaultLang = this.i18n.defaultLang;
    return this.i18n.loadLangData(defaultLang).pipe(
      mergeMap((langData: NzSafeAny) => {
        this.i18n.use(defaultLang, langData);
        return this.viaMock();
      })
    );
  }

  private viaMock(): Observable<void> {
    if (this.tokenService.get()?.token) {
      return this.userService.user().pipe(
        map((response: any) => {
          const user = response.body;
          this.settingService.setUser({
            name: user.username,
            email: user.email,
            avatar: user.avatar
          });
          this.settingService.setApp({
            name: '学生管理系统',
            description: '这是一个学生管理系统'
          });
          this.aclService.setFull(true);
          this.menuService.add([
            {
              text: '查询',
              group: true,
              children: [
                {
                  text: '学生查询',
                  link: '/student/query',
                  icon: { type: 'icon', value: 'book', theme: 'outline' }
                }
              ]
            },
            {
              text: '基础信息管理',
              group: true,
              children: [
                {
                  text: '用户列表',
                  link: '/user/list',

                  icon: { type: 'icon', value: 'book', theme: 'outline' }
                },
                {
                  text: '角色列表',
                  link: '/role/list',

                  icon: { type: 'icon', value: 'book', theme: 'outline' }
                }
              ]
            },
            {
              text: '学工信息管理',
              group: true,
              children: [
                {
                  text: '校区管理',
                  link: '/campus/list',

                  icon: { type: 'icon', value: 'book', theme: 'outline' }
                },
                {
                  text: '系部管理',
                  link: '/department/list',

                  icon: { type: 'icon', value: 'book', theme: 'outline' }
                },
                {
                  text: '专业管理',
                  link: '/major/list',

                  icon: { type: 'icon', value: 'book', theme: 'outline' }
                },
                {
                  text: '班级管理',
                  link: '/class/list',
                  icon: { type: 'icon', value: 'book', theme: 'outline' }
                },
                {
                  text: '学生管理',
                  link: '/student/list',
                  icon: { type: 'icon', value: 'book', theme: 'outline' }
                }
              ]
            }
          ]);
          this.titleService.suffix = '学生管理系统';
        })
      );
    } else {
      return of(void 0);
    }
  }

  load(): Observable<void> {
    return this.viaMockI18n();
  }
}
