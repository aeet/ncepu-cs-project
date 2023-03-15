import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { CampusListComponent } from './list/list.component';

const routes: Routes = [

  { path: 'list', component: CampusListComponent }];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class CampusRoutingModule { }
