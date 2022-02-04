import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import {FormsModule, ReactiveFormsModule} from '@angular/forms';
import { HttpClientModule } from '@angular/common/http';

import { MaterialModule } from './material/material.module';
import { BrowserAnimationsModule }
       from '@angular/platform-browser/animations';

import { AppComponent } from './app.component';
import { ToolbarComponent } from './toolbar/toolbar.component';
import { NavbarComponent } from './navbar/navbar.component';
import { LogobarComponent } from './logobar/logobar.component';
import { AdvertisementsComponent } from './advertisements/advertisements.component';
import { EventCalenderComponent } from './event-calender/event-calender.component';
import { BestofTgComponent } from './bestof-tg/bestof-tg.component';
import { ViewpostComponent } from './viewpost/viewpost.component';
import { NgbModule } from '@ng-bootstrap/ng-bootstrap';
import {FormpostComponent} from "./formpost/formpost.component";
import {MatInputModule} from "@angular/material/input";
import {MatFormFieldModule} from "@angular/material/form-field";
import {MatSelectModule} from "@angular/material/select";
import {MatCheckboxModule} from "@angular/material/checkbox";
import {MatCardModule} from "@angular/material/card";
import {MatNativeDateModule} from "@angular/material/core";
import {MatRadioModule} from "@angular/material/radio";
import {MatButtonModule} from "@angular/material/button";
import {MatDatepickerModule} from "@angular/material/datepicker";

@NgModule({

declarations: [

AppComponent,
  FormpostComponent,
  ToolbarComponent,
 NavbarComponent,
 LogobarComponent,
 AdvertisementsComponent,
 EventCalenderComponent,
 BestofTgComponent,
 ViewpostComponent

],

imports: [

BrowserModule,
FormsModule,
HttpClientModule,
MaterialModule,
BrowserAnimationsModule,
NgbModule,
  MatInputModule,
  MatFormFieldModule,
  MatSelectModule,
  ReactiveFormsModule,
  MatRadioModule,
  MatCardModule,
  MatDatepickerModule,
  MatNativeDateModule,
  MatButtonModule,
  MatCheckboxModule,

],

providers: [],

bootstrap: [AppComponent]

})

export class AppModule { }
