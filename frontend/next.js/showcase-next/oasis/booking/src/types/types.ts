export type Cabin = {
  id: number;
  name: string;
  maxCapacity: number;
  regularPrice: number;
  discount: number;
  image: string;
};

export type Booking = {
  id: number;
  cabinId: number;
  checkIn: string;
  checkOut: string;
  guests: number;
  total: number;
  userId: number;
};
