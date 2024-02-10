"use client";

import { useForm } from "react-hook-form";
import {
  useGMOMultiPayment,
  useGMOMultiPaymentGetTokenResult,
} from "./useGMOMultiPayment";

export type GMOMultiPaymentInput = {
  cardno: number;
  expireMonth: number;
  expireYear: number;
  holdername: string;
  securitycode: number;
};

export const useCreatePaymentMethod = ({
  onCompleted,
  onError,
}: Partial<{
  onCompleted: (data: useGMOMultiPaymentGetTokenResult) => void;
  onError: () => void;
}> = {}) => {
  const form = useForm<GMOMultiPaymentInput>();
  const { getToken } = useGMOMultiPayment();

  const onCreatePaymentMethod = form.handleSubmit(
    async ({ cardno, expireMonth, expireYear, holdername, securitycode }) => {
      const paddingMonth = String(expireMonth).padStart(2, "0");

      try {
        const data = await getToken({
          cardno,
          expire: `${expireYear}${paddingMonth}`,
          holdername,
          securitycode,
        });

        if (data.resultCode === "000") {
          console.log(data.tokenObject.token);
          onCompleted?.(data);
        } else {
          console.log("error");
          onError?.();
        }
      } catch (e) {
        console.log("error");
        onError?.();
      }
    }
  );

  return {
    form,
    onCreatePaymentMethod,
  };
};

type useCreatePaymentMethodType = typeof useCreatePaymentMethod;
export type useCreatePaymentMethodResult =
  ReturnType<useCreatePaymentMethodType>;
// export type useCreatePaymentMethodParameter = Parameters<useCreatePaymentMethodType>[0]
