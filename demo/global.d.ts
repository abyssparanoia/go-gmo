declare global {
  interface Window {
    Multipayment: {
      getToken: (creditCard: {
        cardno: number;
        expire: number; // YYYYMM or YYMM
        holdername: string;
        securitycode: number;
      }) => {
        resultCode: `000` | number;
        tokenObject: {
          isSecurityCodeSet: boolean | string;
          maskedCardNo: string;
          toBeExpiredAt: string;
          token: string;
        };
      };
      init: (apiKey: string) => void;
    };
  }
}
const Multipayment = window.Multipayment;
