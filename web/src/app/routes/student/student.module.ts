import { NgModule, Type } from '@angular/core';
import { SharedModule } from '@shared';

import { StudentListComponent } from './student-list.component';
import { StudentQueryComponent } from './student-query.component';
import { StudentRoutingModule } from './student-routing.module';

const COMPONENTS: Array<Type<void>> = [StudentListComponent, StudentQueryComponent];

@NgModule({
  imports: [SharedModule, StudentRoutingModule],
  declarations: COMPONENTS
})
export class StudentModule {}
