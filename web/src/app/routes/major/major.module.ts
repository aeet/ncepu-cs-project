import { NgModule, Type } from '@angular/core';
import { SharedModule } from '@shared';
import { MajorRoutingModule } from './major-routing.module';
import { MajorListComponent } from './list/list.component';

const COMPONENTS: Type<void>[] = [
  MajorListComponent];

@NgModule({
  imports: [
    SharedModule,
    MajorRoutingModule
  ],
  declarations: COMPONENTS,
})
export class MajorModule { }
