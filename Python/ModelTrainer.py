import os
import psycopg2
import pandas as pd
import torch
import torch.nn as nn
import torch.optim as optim
import joblib
from sklearn.preprocessing import LabelEncoder
from sklearn.model_selection import train_test_split
from sklearn.preprocessing import StandardScaler
from sklearn.metrics import accuracy_score, precision_score, recall_score


# Database connection info - use container name instead of localhost
DB_HOST = os.getenv("DB_HOST", "postgres")  # Use the service name from docker-compose
DB_PORT = os.getenv("DB_PORT", "5432")
DB_NAME = os.getenv("DB_NAME", "gees_db")
DB_USER = os.getenv("DB_USER", "MySecretUser")
DB_PASS = os.getenv("DB_PASS", "MySecretPassword123")

# SQL query to get data
QUERY = "SELECT * FROM feature_data WHERE gestureid IS NOT NULL"

def load_data():
    conn = psycopg2.connect(
        host=DB_HOST,
        port=DB_PORT,
        dbname=DB_NAME,
        user=DB_USER,
        password=DB_PASS
    )
    df = pd.read_sql(QUERY, conn)
    conn.close()
    return df

class GestureModel(nn.Module):
    def __init__(self, input_dim, output_dim):
        super(GestureModel, self).__init__()
        self.net = nn.Sequential(
            nn.Linear(input_dim, 128),
            nn.ReLU(),
            nn.Linear(128, 64),
            nn.ReLU(),
            nn.Linear(64, output_dim)
        )

    def forward(self, x):
        return self.net(x)

def main():

    df = load_data()
    df = df.drop(columns=['id']) 

    X = df.drop(columns=['gestureid']).values
    label_encoder = LabelEncoder()
    y = label_encoder.fit_transform(df['gestureid'].values)
    joblib.dump(label_encoder, "gesture_label_encoder.pkl")

    scaler = StandardScaler()
    X_scaled = scaler.fit_transform(X)
    joblib.dump(scaler, "scaler.pkl")

  
    X_train_np, X_val_np, y_train_np, y_val_np = train_test_split(
        X_scaled, y, test_size=0.5 # Split for training and validation
    )

    X_train = torch.tensor(X_train_np, dtype=torch.float32)
    X_val = torch.tensor(X_val_np, dtype=torch.float32)
    y_train = torch.tensor(y_train_np, dtype=torch.long)
    y_val = torch.tensor(y_val_np, dtype=torch.long)

    input_dim = X_train.shape[1]
    output_dim = len(label_encoder.classes_) 
    model = GestureModel(input_dim, output_dim)

    criterion = nn.CrossEntropyLoss()
    optimizer = optim.Adam(model.parameters(), lr=0.001)

    epochs = 1000
    for epoch in range(epochs):
        model.train()
        optimizer.zero_grad()
        outputs = model(X_train)
        loss = criterion(outputs, y_train)
        loss.backward()
        optimizer.step()

        # Validation phase
        model.eval()
        with torch.no_grad():
            val_outputs = model(X_val)
            _, preds = torch.max(val_outputs, 1)

            acc = accuracy_score(y_val.numpy(), preds.numpy())
            precision = precision_score(y_val.numpy(), preds.numpy(), average='weighted', zero_division=1)
            recall = recall_score(y_val.numpy(), preds.numpy(), average='weighted', zero_division=1)

        print(f"Epoch {epoch+1}/{epochs} - Loss: {loss.item():.4f} - "
              f"Val Acc: {acc:.4f} - Val Precision: {precision:.4f} - Val Recall: {recall:.4f}")

    torch.save(model.state_dict(), "gesture_model.pt")
    print("Model saved as gesture_model.pt")

if __name__ == "__main__":
    main()
