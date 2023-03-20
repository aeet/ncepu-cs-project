import { NgModule, Type } from '@angular/core';
import { SharedModule } from '@shared';

import { StudentRoutingModule } from './student-routing.module';

const COMPONENTS: Array<Type<void>> = [];

@NgModule({
  imports: [SharedModule, StudentRoutingModule],
  declarations: COMPONENTS
})
export class StudentModule {}
