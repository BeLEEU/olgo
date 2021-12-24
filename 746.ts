


function calculateCost(index: number, cost: number[], sum: number[]) {

  if(index == 1) {
    sum[index] = 0;
    return;
  }
  calculateCost(index-1, cost, sum);
  sum[index] = Math.min((sum[index-1] + cost[index-1]), (sum[index-2] + cost[index-2]))
  
  return
}

function minCostClimbingStairs(cost: number[]): number {
  
  let sum: number[] = new Array(cost.length)
  sum[0] = 0;
  sum[1] = 0;
  
  calculateCost(cost.length, cost, sum);
  return Math.min(...sum)
}
console.log(minCostClimbingStairs([1,100,1,1,1,100,1,1,100,1]))
