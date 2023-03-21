import { NgModule, Type } from '@angular/core';
import { SharedModule } from '@shared';

import { StudentListComponent } from './student-list.component';
import { StudentRoutingModule } from './student-routing.module';

const COMPONENTS: Array<Type<void>> = [StudentListComponent];

@NgModule({
  imports: [SharedModule, StudentRoutingModule],
  declarations: COMPONENTS
})
export class StudentModule {}
