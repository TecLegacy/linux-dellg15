import type { Metadata } from 'next';
import { Josefin_Sans } from 'next/font/google';
import './globals.css';
import Header from '@/components/Header';

const josefinSans = Josefin_Sans({ subsets: ['latin'], display: 'swap' });

export const metadata: Metadata = {
  title: {
    template: '%s | Cabin Booking',
    default: 'Cabin Booking',
  },
  description:
    'Book your perfect woodland retreat with our Cabin Booking app. Discover a variety of charming cabins in the woods, from rustic to luxurious, and enjoy seamless booking, detailed descriptions, and real user reviews. Escape to nature today with the ultimate cabin getaway experience.',
};
0;

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang='en'>
      <body
        className={`${josefinSans.className} antialiased bg-primary-950 text-primary-100 min-h-screen flex flex-col relative`}
      >
        <Header />

        <div className='flex-1 px-8 py-12 grid'>
          <main className='max-w-7xl mx-auto w-full'>{children}</main>
        </div>
      </body>
    </html>
  );
}
