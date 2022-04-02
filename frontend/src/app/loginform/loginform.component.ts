import {Component, OnInit} from '@angular/core';
import {ModalDismissReasons, NgbModal} from '@ng-bootstrap/ng-bootstrap';
import {FormControl, FormGroup, Validators} from '@angular/forms';
import {Login, LoginService} from "../services/login.service";

@Component({
  selector: 'app-loginform',
  templateUrl: './loginform.component.html',
  styleUrls: ['./loginform.component.css']
})
export class LoginformComponent implements OnInit {
  closeResult = '';
  isLoginFailed = false
  errorMessage = ''
  isFormSubmitted = false

  constructor(private modalService: NgbModal, private loginService: LoginService) {
  }

  ngOnInit(): void {
  }

  form = new FormGroup({
    password: new FormControl('', [Validators.required, Validators.minLength(6)]),
    email: new FormControl('', [Validators.required, Validators.email])
  });

  get f() {
    return this.form.controls;
  }

  submit() {
    this.isFormSubmitted = true
    console.log(this.form.value);
    const login: Login = {
      Email: this.form.value.email,
      Password: this.form.value.password
    }
    this.loginService.login(login).subscribe(data => {
      console.log(data)
      const status = data["status"]
      this.isLoginFailed = status == 404;
      this.errorMessage = data["message"]
      this.form.reset()
      if (!this.isLoginFailed) {
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

