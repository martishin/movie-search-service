import { SplitFactory } from "@splitsoftware/splitio";

const sdk: SplitIO.IBrowserSDK = SplitFactory({
  core: {
    authorizationKey: import.meta.env.VITE_SPLIT_BROWSER_KEY!,
    key: "asmartishin",
  },
});

export const splitClient: SplitIO.IBrowserClient = sdk.client();
