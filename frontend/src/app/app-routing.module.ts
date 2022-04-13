import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { GridComponent } from './grid/grid.component';
import {LoginformComponent} from "./loginform/loginform.component";
import {FormpostComponent} from "./formpost/formpost.component";

const routes: Routes = [
  {path: 'view/subcategory/:id', component:GridComponent, runGuardsAndResolvers: 'always'},
  {path: '', component:GridComponent},
  {path:'login', component:LoginformComponent},
  {path:'createpost', component:FormpostComponent}
];

@NgModule({
  imports: [RouterModule.forRoot(routes, {
    onSameUrlNavigation: 'reload'
  })],
  exports: [RouterModule]
})
export class AppRoutingModule { }
