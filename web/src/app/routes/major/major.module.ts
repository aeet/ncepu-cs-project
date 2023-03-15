import { NgModule, Type } from '@angular/core';
import { SharedModule } from '@shared';

import { MajorListComponent } from './list/list.component';
import { MajorRoutingModule } from './major-routing.module';

const COMPONENTS: Array<Type<void>> = [MajorListComponent];

@NgModule({
  imports: [SharedModule, MajorRoutingModule],
  declarations: COMPONENTS
})
export class MajorModule {}
