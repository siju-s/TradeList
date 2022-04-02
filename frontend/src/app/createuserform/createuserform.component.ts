import { Component, OnInit } from '@angular/core';
import {NgbModal, ModalDismissReasons} from '@ng-bootstrap/ng-bootstrap';
import { FormGroup, FormControl, Validators} from '@angular/forms';
import {SignupService, User} from "../services/signup.service";
@Component({
  selector: 'app-createuserform',
  templateUrl: './createuserform.component.html',
  styleUrls: ['./createuserform.component.css']
})
export class CreateuserformComponent implements OnInit {

  closeResult = '';
  isSignupFailed = false
  errorMessage = ''
  isFormSubmitted = false

  constructor(private modalService: NgbModal, private signupService: SignupService) {}

  ngOnInit(): void {
  }

  form = new FormGroup({
    firstname: new FormControl('', [Validators.required, Validators.minLength(3)]),
    lastname: new FormControl('', [Validators.required, Validators.minLength(3)]),
    password: new FormControl('', [Validators.required, Validators.minLength(8)]),
    email: new FormControl('', [Validators.required, Validators.email])
  });

  get f(){
    return this.form.controls;
  }

  submit(){
    console.log(this.form.value);
    const user: User = {
      Contact: {
        FirstName: this.form.value.firstname,
        LastName: this.form.value.lastname,
        Email: this.form.value.email,
        Password: this.form.value.password
      }
    }
    this.signupService.signup(user).subscribe(data => {
      console.log(data)
      const status = data["status"]
      this.isSignupFailed = status == 0;
      this.errorMessage = data["message"]
      this.form.reset()
      if (!this.isSignupFailed) {
        this.modalService.dismissAll()
      }
    })
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
