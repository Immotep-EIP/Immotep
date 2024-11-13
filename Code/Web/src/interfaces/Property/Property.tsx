export interface Tenant {
  id: number;
  name: string;
  email: string;
}

export interface Photo {
  id: number;
  url: string;
}

export interface Damage {
  id: number;
  read: boolean;
  date: string;
  comment: string;
  interventionDate: string;
  room: string;
  priority: string;
  photos: Photo[];
}

export interface Document {
  id: number;
  name: string;
  url: string;
}

export interface Furniture {
  id: number;
  name: string;
  quantity: number;
}

export interface Inventory {
  room: string;
  furniture: Furniture[];
}

export interface RealProperty {
  id: number;
  adress: string;
  city: string;
  zipCode: string;
  image: string;
  status: string;
  tenants: Tenant[];
  startDate: string;
  endDate: string;
  rent: number;
  deposit: number;
  area: number;
  damages: Damage[];
  documents: Document[];
  inventory: Inventory[];
}
