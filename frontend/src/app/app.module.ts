import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { FormsModule } from '@angular/forms';
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

@NgModule({

declarations: [

AppComponent,
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
NgbModule 

],

providers: [],

bootstrap: [AppComponent]

})

export class AppModule { }