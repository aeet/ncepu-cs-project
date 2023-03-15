import { NgModule, Type } from '@angular/core';
import { SharedModule } from '@shared';
import { DepartmentRoutingModule } from './department-routing.module';
import { DepartmentListComponent } from './list/list.component';

const COMPONENTS: Type<void>[] = [
  DepartmentListComponent];

@NgModule({
  imports: [
    SharedModule,
    DepartmentRoutingModule
  ],
  declarations: COMPONENTS,
})
export class DepartmentModule { }
