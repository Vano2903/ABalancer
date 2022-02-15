<?php
#starting the timer
$time_start = hrtime(true);

#connect to mariadb using pdo
try {
    $db = new PDO('mysql:host=db;dbname=benchmarks', 'root', 'root');
} catch (PDOException $e) {
    print "Error!: " . $e->getMessage() . "<br/>";
    die();
}

//generate a random number between 1 and 10
$rand = rand(1, 10);
if ($rand == 1){
    $useless = 0;
    for ($i = 0; $i < 100000000; $i++) {
        $useless += $i;
    }
}

//ending the timer and save on db
$time_end = hrtime(true);
$time = $time_end - $time_start;

//insert into the database
$stmt = $db->prepare("INSERT INTO `benchmarks` (`id`, `backType`, `execTime`) VALUES (NULL, 'php', $time)");
$stmt->execute();

echo ("php ".$time." ns");

?>
