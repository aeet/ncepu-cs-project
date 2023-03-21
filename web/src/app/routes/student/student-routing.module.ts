import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';

import { StudentListComponent } from './student-list.component';

const routes: Routes = [{ path: 'list', component: StudentListComponent }];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class StudentRoutingModule {}
