import { NgModule, Type } from '@angular/core';
import { SharedModule } from '@shared';
import { CampusRoutingModule } from './campus-routing.module';
import { CampusListComponent } from './list/list.component';

const COMPONENTS: Type<void>[] = [
  CampusListComponent];

@NgModule({
  imports: [
    SharedModule,
    CampusRoutingModule
  ],
  declarations: COMPONENTS,
})
export class CampusModule { }
