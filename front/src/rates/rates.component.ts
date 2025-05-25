import {Component, inject, Injectable, Input, model, OnInit} from '@angular/core';
import {HttpClient} from "@angular/common/http";
import {
  MatActionList,
  MatList,
  MatListItem,
  MatListItemLine,
  MatListItemTitle,
} from '@angular/material/list';
import {KeyValuePipe} from '@angular/common';
import {
  MAT_DIALOG_DATA,
  MatDialog,
  MatDialogActions,
  MatDialogClose,
  MatDialogContent,
  MatDialogRef,
  MatDialogTitle
} from '@angular/material/dialog';
import {MatInput, MatInputModule} from '@angular/material/input';
import {MatButton, MatButtonModule} from '@angular/material/button';
import {MatFormFieldModule} from '@angular/material/form-field';
import {FormsModule} from '@angular/forms';

@Injectable({providedIn: 'root'})

@Component({
  selector: 'app-rates',
  templateUrl: './rates.component.html',
  standalone: true,
  styleUrl: './rates.component.css',
  imports: [
    KeyValuePipe,
    MatList,
    MatListItem,
    MatListItemTitle,
    MatListItemLine,
    MatActionList
  ]
})
export class RatesComponent implements OnInit {
  private http = inject(HttpClient);
  public rates: Map<string, Bank[]> = new Map<string, Bank[]>([]).set("", []);

  readonly dialog = inject(MatDialog);

  @Input() newRates!: string;

  public ngOnInit() {
    this.http.get<Rates>('http://localhost:8080/api/rates/param').subscribe(rates => {
      this.rates = new Map(Object.entries(rates.rates));
    })
  }

  deleteBank(key: string, index: number) {
    let banks = this.rates.get(key);
    if (banks) {
      delete banks[index];
      this.rates.set(key, banks);
    }
  }

  editBank(key: string, index: number) {
    let banks = this.rates.get(key)
    let bank: Bank | null
    if (banks) {
      bank = banks[index]
    } else {
      return
    }

    const dialogRef = this.dialog.open(EditBankDialog, {
      data: bank,
    })

    dialogRef.afterClosed().subscribe(result => {
      if (result) {
        let newBank: Bank = result
        let banks = this.rates.get(key);
        if (banks) {
          banks[index] = newBank;
          this.rates.set(key, banks);
        }
      }
    })
  }

  editRate(key: string) {
    const dialogRef = this.dialog.open(EditRateNameDialog, {
      data: key,
    })

    dialogRef.afterClosed().subscribe(result => {
      if (result) {
        let banks = this.rates.get(key)
        if (banks) {
          this.rates.set(result, banks)
          this.rates.delete(key)
        }
      }
    })
  }

  deleteRate(key: string) {
    this.rates.delete(key);
  }

  save() {
    this.http.patch('http://localhost:8080/api/rates/param', {"rates": this.rates}, { observe: 'response' , responseType: 'text'}).subscribe(response => {
      if (response.status != 200) {
        console.error('Error occurred while saving rates.');
      }
    })
  }

  addRate() {
    const dialogRef = this.dialog.open(AddRateNameDialog)

    dialogRef.afterClosed().subscribe(result => {
      if (result) {
        if (!this.rates.has(result)) {
          this.rates.set(result, [])
        }
      }
    })
  }

  addBank(key: string) {
    const dialogRef = this.dialog.open(AddBankDialog)

    dialogRef.afterClosed().subscribe(result => {
      if (result) {
        let newBank: Bank = result
        let banks = this.rates.get(key);
        if (banks) {
          banks.push(newBank);
          this.rates.set(key, banks);
        }
      }
    })
  }
}

export interface Rates {
  rates: Map<string, Bank[]>;
}

export interface Bank {
  bank_name: string;
  rate_value: number;
}

@Component({
  selector: 'edit-bank-dialog',
  template: `
    <mat-dialog-content>
      <mat-form-field>
        <mat-label>New bank name</mat-label>
        <input matInput [(ngModel)]="newBankName" />
      </mat-form-field>
      <mat-form-field>
        <mat-label>New rate value</mat-label>
        <input matInput [(ngModel)]="newRate" />
      </mat-form-field>
    </mat-dialog-content>
    <mat-dialog-actions>
      <button mat-button (click)="close()">Close</button>
      <button mat-button [mat-dialog-close]="returnNewBank()" cdkFocusInitial>Save</button>
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
export class EditBankDialog {
  readonly dialogRef = inject(MatDialogRef<EditBankDialog>);
  readonly oldBank = inject<Bank>(MAT_DIALOG_DATA);

  newBankName: string = this.oldBank.bank_name
  newRate: number = this.oldBank.rate_value

  close(): void {
    this.dialogRef.close();
  }

  returnNewBank(): Bank {
    return {"bank_name": this.newBankName, "rate_value": this.newRate}
  }
}

@Component({
  selector: 'edit-rate-dialog',
  template: `
    <mat-dialog-content>
      <mat-form-field>
        <mat-label>New rate name</mat-label>
        <input matInput [(ngModel)]="newRateName" />
      </mat-form-field>
    </mat-dialog-content>
    <mat-dialog-actions>
      <button mat-button (click)="close()">Close</button>
      <button mat-button [mat-dialog-close]="returnNewRateName()" cdkFocusInitial>Save</button>
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
export class EditRateNameDialog {
  readonly dialogRef = inject(MatDialogRef<EditRateNameDialog>);
  readonly oldRateName: string = inject<string>(MAT_DIALOG_DATA);

  newRateName: string = this.oldRateName

  close(): void {
    this.dialogRef.close();
  }

  returnNewRateName(): string {
    return this.newRateName
  }
}

@Component({
  selector: 'add-rate-dialog',
  template: `
    <mat-dialog-content>
      <mat-form-field>
        <mat-label>New rate name</mat-label>
        <input matInput [(ngModel)]="newRateName" />
      </mat-form-field>
    </mat-dialog-content>
    <mat-dialog-actions>
      <button mat-button (click)="close()">Close</button>
      <button mat-button [mat-dialog-close]="returnNewRateName()" cdkFocusInitial>Save</button>
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
export class AddRateNameDialog {
  readonly dialogRef = inject(MatDialogRef<AddRateNameDialog>);
  newRateName: string = ""

  close(): void {
    this.dialogRef.close();
  }

  returnNewRateName(): string {
    return this.newRateName
  }
}

@Component({
  selector: 'add-bank-dialog',
  template: `
    <mat-dialog-content>
      <mat-form-field>
        <mat-label>New bank name</mat-label>
        <input matInput [(ngModel)]="newBankName" />
      </mat-form-field>
      <mat-form-field>
        <mat-label>New rate value</mat-label>
        <input matInput [(ngModel)]="newRate" />
      </mat-form-field>
    </mat-dialog-content>
    <mat-dialog-actions>
      <button mat-button (click)="close()">Close</button>
      <button mat-button [mat-dialog-close]="returnNewBank()" cdkFocusInitial>Save</button>
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
export class AddBankDialog {
  readonly dialogRef = inject(MatDialogRef<AddBankDialog>);

  newBankName: string = ""
  newRate: number = 0

  close(): void {
    this.dialogRef.close();
  }

  returnNewBank(): Bank {
    return {"bank_name": this.newBankName, "rate_value": this.newRate}
  }
}

