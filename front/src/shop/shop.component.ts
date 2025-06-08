import {Component, inject, OnInit} from '@angular/core';
import {FormsModule} from '@angular/forms';
import {HttpClient} from '@angular/common/http';
import {
  MatDialog,
  MatDialogActions,
  MatDialogClose,
  MatDialogContent,
  MatDialogRef,
  MatDialogTitle
} from '@angular/material/dialog';
import {MatFormFieldModule} from '@angular/material/form-field';
import {MatInput, MatInputModule} from '@angular/material/input';
import {MatButton, MatButtonModule} from '@angular/material/button';
import {KeyValuePipe} from '@angular/common';
import {MatActionList, MatList, MatListItem, MatListItemLine, MatListItemTitle} from '@angular/material/list';

@Component({
  selector: 'app-shop',
  imports: [
    KeyValuePipe,
    MatActionList,
    MatList,
    MatListItem,
    MatListItemLine,
    MatListItemTitle
  ],
  templateUrl: './shop.component.html',
  standalone: true,
  styleUrl: './shop.component.css'
})
export class ShopComponent implements OnInit  {
  private http = inject(HttpClient);
  public transactions: Transaction[] = []
  readonly dialog = inject(MatDialog);

  public ngOnInit() {
    this.refreshTransactions()
  }

  pay() {
    const dialogRef = this.dialog.open(PayDialog)

    dialogRef.afterClosed().subscribe(result => {
      if (result) {
        if (result.isNotEmpty()) {
          this.http.post('http://localhost:8080/api/gateway/pay', result, { observe: 'response'}).subscribe(response => {
            if (response.status == 200) {
              this.refreshTransactions()
            }
          })
        }
      }
    })
  }

  updateStatus(id: number, index: number) {
    this.http.get('http://localhost:8080/api/gateway/pay_status/'+id, { observe: 'response', responseType: 'text'}).subscribe(response => {
      if (response.status == 200) {
        this.transactions[index].status = Number(response.body);
      }
    })
  }

  refreshTransactions() {
    this.http.get<Transaction[]>('http://localhost:8080/api/transactions/all').subscribe(transactions => {
      this.transactions = transactions.sort((a, b) => {
        if(a.id > b.id) {
          return -1;
        } else if(a.id < b.id) {
          return 1;
        } else {
          return 0;
        }
      });
    })
  }

}

@Component({
  selector: 'pay-dialog',
  template: `
    <mat-dialog-content>
      <mat-form-field>
        <mat-label>Сумма</mat-label>
        <input matInput [(ngModel)]="data.amount" />
      </mat-form-field>
      <mat-form-field>
        <mat-label>Код валюты</mat-label>
        <input matInput [(ngModel)]="data.currency_code" />
      </mat-form-field>
      <mat-form-field>
        <mat-label>PAN карты</mat-label>
        <input matInput [(ngModel)]="data.pan" />
      </mat-form-field>
      <mat-form-field>
        <mat-label>Expired карты</mat-label>
        <input matInput [(ngModel)]="data.expired" />
      </mat-form-field>
      <mat-form-field>
        <mat-label>CVV карты</mat-label>
        <input matInput [(ngModel)]="data.cvv" />
      </mat-form-field>
    </mat-dialog-content>
    <mat-dialog-actions>
      <button mat-button (click)="close()">Закрыть</button>
      <button mat-button [mat-dialog-close]="returnData()" cdkFocusInitial>Оплатить</button>
    </mat-dialog-actions>
  `,
  imports: [
    MatFormFieldModule,
    MatInputModule,
    FormsModule,
    MatButtonModule,
    MatDialogTitle,
    MatDialogContent,
    MatDialogActions,
    MatDialogClose,
    MatInput,
    MatButton,
  ],
  standalone: true
})
export class PayDialog {
  readonly dialogRef = inject(MatDialogRef<PayDialog>);

  data: PaymentData = new PaymentData(0, 0, "", "", "")

  close(): void {
    this.dialogRef.close();
  }

  returnData(): PaymentData {
    this.data.amount = Number(this.data.amount);
    this.data.currency_code = Number(this.data.currency_code);
    return this.data;
  }
}

export interface Transaction {
  id: number;
  currency_code: number;
  amount: number;
  bank_name: string;
  status: number;
  created_at: string;
}

export class PaymentData {
  constructor(
    public amount: number,
    public currency_code: number,
    public pan: string,
    public expired: string,
    public cvv: string,
  ) {}
  public isNotEmpty(): boolean {
    return (
      this.amount != 0 && this.currency_code != 0 && this.pan != "" && this.expired != "" && this.cvv != ""
    );
  }
}

