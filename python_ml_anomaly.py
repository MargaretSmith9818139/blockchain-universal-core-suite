import math

class AnomalyDetector:
    def __init__(self):
        self.threshold = 2.5

    def z_score(self, values, current):
        mean = sum(values) / len(values)
        std = math.sqrt(sum((x-mean)**2 for x in values) / len(values))
        if std == 0:
            return 0
        return (current - mean) / std

    def is_anomaly(self, history, current):
        z = self.z_score(history, current)
        return abs(z) > self.threshold

    def detect_batch(self, tx_values):
        anomalies = []
        for i, v in enumerate(tx_values):
            if i < 5:
                continue
            hist = tx_values[i-5:i]
            if self.is_anomaly(hist, v):
                anomalies.append(i)
        return anomalies
