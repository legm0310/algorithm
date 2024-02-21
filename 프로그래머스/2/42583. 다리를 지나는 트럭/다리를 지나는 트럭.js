function solution(bridge_length, weight, truck_weights) {
  let sec = 0,
    weightSum = 0;
  const onBridge = [];
  const truckSet = truck_weights.reduce((acc, value) => {
    acc.push([value, bridge_length]);
    return acc;
  }, []);

  while (true) {
    sec++;
    onBridge.forEach((v) => {
      if (v) v[1]--;
    });
      
    //지나간 트럭 먼저 제거(다리의 최대무게와 현재 총 트럭무게 비교)
    if (onBridge[0]?.[1] <= 0) {
      weightSum -= onBridge[0]?.[0];
      onBridge.shift();
    }
    //트럭 --> 다리
    if (weightSum + truckSet[0]?.[0] <= weight) {
      weightSum += truckSet[0]?.[0];
      onBridge.push(truckSet.shift());
    }

    if (truckSet.length === 0 && onBridge.length === 0) break;
  }
  return sec;
}