SELECT house.BUYER_ID
     , SUM(price.PRICE) AS TOTAL_WORTH
  FROM house
  LEFT 
  JOIN price
    ON house.HOUSE_ID = price.HOUSE_ID
 GROUP 
    BY house.BUYER_ID
     , price.PRICE
 HAVING COUNT(house.HOUSE_ID) > 1