import { Component } from '@angular/core';
import {MatTab, MatTabGroup} from '@angular/material/tabs';
import {ShopComponent} from '../shop/shop.component';
import {RatesComponent} from '../rates/rates.component';

@Component({
  selector: 'app-root',
  imports: [ShopComponent, RatesComponent, MatTab, MatTabGroup],
  templateUrl: './app.component.html',
  standalone: true,
  styleUrl: './app.component.css'
})
export class AppComponent {
  title = 'front';
}

