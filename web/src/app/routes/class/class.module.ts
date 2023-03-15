import { NgModule, Type } from '@angular/core';
import { SharedModule } from '@shared';
import { ClassRoutingModule } from './class-routing.module';
import { ClassListComponent } from './list/list.component';

const COMPONENTS: Type<void>[] = [
  ClassListComponent];

@NgModule({
  imports: [
    SharedModule,
    ClassRoutingModule
  ],
  declarations: COMPONENTS,
})
export class ClassModule { }
