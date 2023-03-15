import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { environment } from '@env/environment';
import { SimpleGuard } from '@yelon/auth';

import { LayoutBasicComponent } from '../layout/basic/basic.component';
import { LayoutPassportComponent } from '../layout/passport/passport.component';
import { DashboardComponent } from './dashboard/dashboard.component';
import { UserLockComponent } from './passport/lock/lock.component';
import { UserLoginComponent } from './passport/login/login.component';

const routes: Routes = [
  {
    path: '',
    canActivate: [SimpleGuard],
    component: LayoutBasicComponent,
    children: [
      { path: '', redirectTo: 'dashboard', pathMatch: 'full' },
      { path: 'dashboard', component: DashboardComponent, data: { title: '仪表盘' } },
      { path: 'exception', loadChildren: () => import('./exception/exception.module').then(m => m.ExceptionModule) },
      { path: 'role', loadChildren: () => import('./role/role.module').then(m => m.RoleModule) },
      { path: 'campus', loadChildren: () => import('./campus/campus.module').then(m => m.CampusModule) },
      { path: 'department', loadChildren: () => import('./department/department.module').then(m => m.DepartmentModule) },
      { path: 'major', loadChildren: () => import('./major/major.module').then(m => m.MajorModule) },
      { path: 'class', loadChildren: () => import('./class/class.module').then(m => m.ClassModule) }
    ]
  },
  {
    path: 'passport',
    component: LayoutPassportComponent,
    children: [
      { path: 'login', component: UserLoginComponent, data: { title: '登录' } },
      { path: 'lock', component: UserLockComponent, data: { title: '锁屏' } }
    ]
  },
  { path: '**', redirectTo: 'exception/404' }
];

@NgModule({
  imports: [
    RouterModule.forRoot(routes, {
      useHash: environment.useHash,
      scrollPositionRestoration: 'top'
    })
  ],
  exports: [RouterModule]
})
export class RouteRoutingModule {}
