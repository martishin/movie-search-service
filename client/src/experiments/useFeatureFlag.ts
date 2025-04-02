import { useEffect, useState } from "react";

import { splitClient } from "./splitClient";

export function useFeatureFlag(flagName: string, defaultValue = false): boolean {
  const [enabled, setEnabled] = useState(defaultValue);

  useEffect(() => {
    const treatment = splitClient.getTreatment(flagName);
    if (treatment !== "control") {
      setEnabled(treatment === "on");
    } else {
      const onReady = () => {
        const finalTreatment = splitClient.getTreatment(flagName);
        setEnabled(finalTreatment === "on");
      };

      splitClient.on(splitClient.Event.SDK_READY, onReady);

      return () => {
        splitClient.off(splitClient.Event.SDK_READY, onReady);
      };
    }
  }, [flagName]);

  return enabled;
}
