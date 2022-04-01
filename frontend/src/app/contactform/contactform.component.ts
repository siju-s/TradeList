import { Component, OnInit } from '@angular/core';
import {NgbModal, ModalDismissReasons} from '@ng-bootstrap/ng-bootstrap';
import { FormGroup, FormControl, Validators} from '@angular/forms';
@Component({
  selector: 'app-contactform',
  templateUrl: './contactform.component.html',
  styleUrls: ['./contactform.component.css']
})
export class ContactformComponent implements OnInit {

  closeResult = '';

  constructor(private modalService: NgbModal) {}

  ngOnInit(): void {
  }

  form = new FormGroup({
    fullname: new FormControl('', [Validators.required, Validators.minLength(3)]),
    message: new FormControl('', [Validators.required]),
    email: new FormControl('', [Validators.required, Validators.email])
  });

  get f(){
    return this.form.controls;
  }
  
  submit(){
    console.log(this.form.value);
  }
  open(content: any) {
    this.modalService.open(content, {ariaLabelledBy: 'modal-basic-title'}).result.then((result) => {
      this.closeResult = `${result}`;
    }, (reason) => {
      this.closeResult = `${this.getDismissReason(reason)}`;
    });
  }

  private getDismissReason(reason: any): string {
    if (reason === ModalDismissReasons.ESC) {
      return '';
    } else if (reason === ModalDismissReasons.BACKDROP_CLICK) {
      return '';
    } else {
      return ``;
    }
  }
} 

